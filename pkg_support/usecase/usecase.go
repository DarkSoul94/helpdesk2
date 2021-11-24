package usecase

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/global_const/actions"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

type SupportUsecase struct {
	repo   pkg_support.ISupportRepo
	ticket pkg_ticket.IUCForSupport
	perm   group_manager.IPermManager
}

func NewSupportUsecase(
	repo pkg_support.ISupportRepo,
	perm group_manager.IPermManager,
	ticket pkg_ticket.IUCForSupport,
) *SupportUsecase {
	return &SupportUsecase{
		repo:   repo,
		perm:   perm,
		ticket: ticket,
	}
}

//CreateSupport создание нового саппорта
func (u *SupportUsecase) CreateSupport(usersID ...uint64) models.Err {
	for _, userID := range usersID {
		supp, _ := u.repo.GetSupport(userID)
		if supp != nil {
			continue
		}
		//Создание нового объекта саппорта и внесение его в базу
		supp = internal_models.NewSupport(userID)
		if err := u.repo.CreateSupport(supp); err != nil {
			return err
		}
		//Создание нового объекта карточки саппорта и внесение ее в базу
		card := internal_models.NewSupportCard(userID)
		if err := u.repo.CreateCard(card); err != nil {
			return err
		}
	}
	return nil
}

//DeleteSupport - удаление саппорта из базы
func (u *SupportUsecase) DeleteSupport(usersID ...uint64) models.Err {
	for _, userID := range usersID {
		//проверка что суппорт с таким ID есть в списке, если нет - переходим к следующему
		supp, _ := u.repo.GetSupport(userID)
		if supp == nil {
			continue
		}
		card, err := u.repo.GetCardBySupportID(userID)
		if err != nil {
			return err
		}
		//если у этого суппорта в карточке проставлен признак старшего, то обновляем информацию у его подчиненных
		if card.IsSenior {
			if err := u.repo.ResetSenior(userID); err != nil {
				return err
			}
		}
		if err := u.repo.DeleteCard(userID); err != nil {
			return err
		}
		if err := u.repo.DeleteSupport(userID); err != nil {
			return err
		}
	}
	return nil
}

func (u *SupportUsecase) GetSupportList() ([]*internal_models.Support, models.Err) {
	supports, err := u.repo.GetSupportList()
	if err != nil {
		return nil, err
	}
	return supports, nil
}

func (u *SupportUsecase) GetStatusesList() ([]*internal_models.Status, models.Err) {
	statuses, err := u.repo.GetStatusesList()
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

func (u *SupportUsecase) GetActiveSupports() ([]*internal_models.Support, models.Err) {
	return u.repo.GetActiveSupports()
}

func (u *SupportUsecase) GetSupportForDistribution(supportID uint64) *internal_models.Support {
	var support = new(internal_models.Support)
	if supportID != 0 {
		if u.repo.CheckForActivity(supportID) {
			support, _ = u.repo.GetSupport(supportID)
			return support
		}
	}

	prioritized := u.repo.GetPrioritizedSupportID()
	if u.repo.CheckForBusy(prioritized) {
		support, _ = u.repo.GetRandomFreeSupport()
	} else {
		support, _ = u.repo.GetSupport(prioritized)
	}
	return support
}

func (u *SupportUsecase) AddSupportActivity(support *internal_models.Support, ticketID uint64) models.Err {
	if support.Priority {
		for _, val := range u.priorityHelper(support) {
			if err := u.repo.UpdateSupport(val); err != nil {
				return err
			}
		}
	}
	return u.repo.CreateSupportActivity(support.ID, ticketID)
}

func (u *SupportUsecase) RemoveSupportActivity(ticketID uint64) models.Err {
	return u.repo.RemoveSupportActivity(ticketID)
}

func (u *SupportUsecase) UpdateSupportActivity(supportID, ticketID uint64) models.Err {
	return u.repo.UpdateSupportActivity(supportID, ticketID)
}

func (u *SupportUsecase) SetSupportStatus(supportID, statusID uint64) models.Err {
	var (
		support internal_models.Support = internal_models.Support{
			ID:     supportID,
			Status: &internal_models.Status{},
		}
	)
	shift, err := u.repo.GetLastShift(supportID)
	if shift.ClosingStatus || err != nil {
		return supportErr_ClosedShift
	}

	if support.Status, err = u.repo.GetStatus(statusID); err != nil {
		return err
	}

	for _, val := range u.priorityHelper(&support) {
		if err := u.repo.UpdateSupport(val); err != nil {
			return err
		}
	}
	return u.statusHistoryHelper(&support, shift.ID)
}

func (u *SupportUsecase) OpenShift(supportID uint64, user *models.User) models.Err {
	shift, err := u.repo.GetLastShift(supportID)
	if err != nil {
		shift = new(internal_models.Shift)
		shift.ClosingStatus = true
	}
	if shift.ClosingStatus {
		if shift.WasOpenedToday() {
			if !u.perm.CheckPermission(user.Group.ID, actions.AdminTA) {
				return supportErr_CannotReopen
			}
			shift.Reopen()
			return u.updateShift(shift)
		}
		//TODO добавить проверку на опоздание по графику и можно ли вообще открывать смену
		shift.Open(supportID, time.Now())
		return u.updateShift(shift)
	}
	return supportErr_AlreadyOpen
}

func (u *SupportUsecase) CloseShift(supportID uint64, user *models.User) models.Err {
	shift, err := u.repo.GetLastShift(supportID)
	if err != nil {
		return err
	}
	if !shift.ClosingStatus {
		if u.repo.CheckForBusy(supportID) {
			if !u.perm.CheckPermission(user.Group.ID, actions.AdminTA) {
				return supportErr_Busy
			}
			//TODO добавить возврат запросов на распределение если смену закрывает админ
		}
		shift.Close()
		return u.updateShift(shift)
	}
	return supportErr_ClosedShift
}

func (u *SupportUsecase) updateShift(shift *internal_models.Shift) models.Err {
	var (
		err models.Err
		id  uint64
	)
	if id, err = u.repo.UpdateShift(shift); err != nil {
		return err
	}
	if shift.ID == 0 {
		shift.ID = id
	}
	if shift.Support.Status != nil {
		shift.Support.Status, err = u.repo.GetStatus(shift.Support.Status.ID)
		if err != nil {
			return err
		}
		forUpdate := u.priorityHelper(shift.Support)
		for _, support := range forUpdate {
			if err := u.repo.UpdateSupport(support); err != nil {
				return err
			}
		}
		if err := u.statusHistoryHelper(shift.Support, shift.ID); err != nil {
			return err
		}
	}
	return nil
}

func (u *SupportUsecase) GetLastShift(supportID uint64) (*internal_models.Shift, models.Err) {
	return u.repo.GetLastShift(supportID)
}

func (u *SupportUsecase) GetSupportStatus(supportID uint64) (*internal_models.Status, models.Err) {
	supp, err := u.repo.GetSupport(supportID)
	if err != nil {
		return nil, err
	}
	return u.repo.GetStatus(supp.Status.ID)
}

func (u *SupportUsecase) GetCurrentStatuses() ([]*internal_models.SupportInfo, map[string]int, models.Err) {
	infoArray := make([]*internal_models.SupportInfo, 0)
	suppList, err := u.repo.GetSupportListForToday()
	if err != nil {
		return nil, nil, err
	}
	for _, supp := range suppList {
		info, err := u.getInfoHelper(supp)
		if err != nil {
			return nil, nil, err
		}
		infoArray = append(infoArray, info)
	}
	return infoArray, u.totalInfoHelper(), nil
}

func (u *SupportUsecase) GetCard(cardID uint64) (*internal_models.Card, models.Err) {
	return u.repo.GetCard(cardID)
}

func (u *SupportUsecase) UpdateCard(card *internal_models.Card) models.Err {
	currentCard, err := u.repo.GetCard(card.ID)
	if err != nil {
		return err
	}
	if currentCard.IsSenior && !card.IsSenior {
		if err := u.repo.ResetSenior(card.ID); err != nil {
			return err
		}
	}
	if currentCard.Senior.ID != card.Senior.ID {
		if card.Senior.ID == 0 {
			if err := u.repo.ResetSenior(card.ID); err != nil {
				return err
			}
		} else {
			seniorCard, err := u.repo.GetCardBySupportID(card.Senior.ID)
			if err != nil {
				return err
			}
			card.Color = seniorCard.Color
		}
	}
	return u.repo.UpdateCard(card)
}
