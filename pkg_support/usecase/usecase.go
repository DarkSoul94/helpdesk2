package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
)

type SupportUsecase struct {
	repo pkg_support.ISupportRepo
	perm group_manager.IPermManager
}

func NewSupportUsecase(repo pkg_support.ISupportRepo, perm group_manager.IPermManager) *SupportUsecase {
	return &SupportUsecase{
		repo: repo,
		perm: perm,
	}
}

func (u *SupportUsecase) CreateSupport(usersID ...uint64) models.Err {
	for _, userID := range usersID {
		supp, _ := u.repo.GetSupport(userID)
		if supp != nil {
			continue
		}
		supp = internal_models.NewSupport(userID)
		if err := u.repo.CreateSupport(supp); err != nil {
			return err
		}
		card := internal_models.NewSupportCard(userID)
		if err := u.repo.CreateCard(card); err != nil {
			return err
		}
	}
	return nil
}

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

	forUpdate := u.priorityHelper(&support)
	for _, val := range forUpdate {
		if err := u.repo.UpdateSupport(val); err != nil {
			return err
		}
	}
	return u.statusHistoryHelper(&support, shift.ID)
}
