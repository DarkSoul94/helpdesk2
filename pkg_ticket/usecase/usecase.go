package usecase

import (
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/cachettl"
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	"github.com/spf13/viper"
)

type TicketUsecase struct {
	repo     pkg_ticket.ITicketRepo
	catSecUC cat_sec_manager.ICatSecUsecase
	regFilUC reg_fil_manager.IRegFilUsecase
	permUC   group_manager.IPermManager
	store    cachettl.ObjectStore
}

func NewTicketUsecase(
	tRepo pkg_ticket.ITicketRepo,
	catSecUC cat_sec_manager.ICatSecUsecase,
	regFilUC reg_fil_manager.IRegFilUsecase,
	permUC group_manager.IPermManager,
) *TicketUsecase {
	return &TicketUsecase{
		repo:     tRepo,
		catSecUC: catSecUC,
		regFilUC: regFilUC,
		permUC:   permUC,
		store:    *cachettl.NewObjectStore(viper.GetDuration("app.ttl_cache.clear_period") * time.Second),
	}
}

func (u *TicketUsecase) CreateCategory(cat *internal_models.Category) (uint64, models.Err) {
	if len(cat.Name) == 0 {
		return 0, models.BadRequest("Имя не должно быть пустым")
	}
	return u.catSecUC.CreateCategory(cat)
}

func (u *TicketUsecase) UpdateCategory(cat *internal_models.Category) models.Err {
	if len(cat.Name) == 0 {
		return models.BadRequest("Имя не должно быть пустым")
	}

	return u.catSecUC.UpdateCategory(cat)
}

func (u *TicketUsecase) CreateCategorySection(sec *internal_models.CategorySection) (uint64, models.Err) {
	if sec.CategoryID == 0 {
		return 0, models.BadRequest("Не указана категория")
	}

	if len(sec.Name) == 0 {
		return 0, models.BadRequest("Имя не должно быть пустым")
	}

	if sec.NeedApproval && len(sec.ApprovalGroups) == 0 {
		return 0, models.BadRequest("Не выбрано ни одной согласовывающей группы")
	}

	return u.catSecUC.CreateCategorySection(sec)
}

func (u *TicketUsecase) UpdateCategorySection(sec *internal_models.CategorySection) models.Err {
	if sec.CategoryID == 0 {
		return models.BadRequest("Не указана категория")
	}

	if len(sec.Name) == 0 {
		return models.BadRequest("Имя не должно быть пустым")
	}

	if sec.NeedApproval && len(sec.ApprovalGroups) == 0 {
		return models.BadRequest("Не выбрано ни одной согласовывающей группы")
	}

	return u.catSecUC.UpdateCategorySection(sec)
}

func (u *TicketUsecase) GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, models.Err) {
	return u.catSecUC.GetCategorySection(forSearch)
}

func (u *TicketUsecase) GetCategorySectionList() ([]internal_models.CategorySectionList, models.Err) {
	return u.catSecUC.GetCategorySectionList()
}

func (u *TicketUsecase) CreateRegion(reg *internal_models.Region) (uint64, models.Err) {
	if len(reg.Name) == 0 {
		return 0, models.BadRequest("Имя региона не должно быть пустым")
	}

	return u.regFilUC.CreateRegion(reg)
}

func (u *TicketUsecase) UpdateRegion(reg *internal_models.Region) models.Err {
	if len(reg.Name) == 0 {
		return models.BadRequest("Имя региона не должно быть пустым")
	}

	return u.regFilUC.UpdateRegion(reg)
}

func (u *TicketUsecase) DeleteRegion(id uint64) models.Err {
	return u.regFilUC.DeleteRegion(id)
}

func (u *TicketUsecase) CreateFilial(fil *internal_models.Filial) (uint64, models.Err) {
	if fil.RegionID == 0 {
		return 0, models.BadRequest("У данного филиала не указан регион")
	}

	if len(fil.Name) == 0 {
		return 0, models.BadRequest("Имя не должно быть пустым")
	}

	return u.regFilUC.CreateFilial(fil)
}

func (u *TicketUsecase) UpdateFilial(fil *internal_models.Filial) models.Err {
	if fil.RegionID == 0 {
		return models.BadRequest("У данного филиала не указан регион")
	}

	if len(fil.Name) == 0 {
		return models.BadRequest("Имя не должно быть пустым")
	}

	return u.regFilUC.UpdateFilial(fil)
}

func (u *TicketUsecase) DeleteFilial(id uint64) models.Err {
	return u.regFilUC.DeleteFilial(id)
}

func (u *TicketUsecase) GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, models.Err) {
	return u.regFilUC.GetRegionsWithFilials()
}

func (u *TicketUsecase) GetTicketStatuses(group_id uint64, all bool) ([]*internal_models.TicketStatus, models.Err) {
	list, err := u.repo.GetTicketStatuses()
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	if !all {
		list = u.ticketStatusFilter(list, group_id)
	}

	return list, nil
}

func (u *TicketUsecase) ticketStatusFilter(list []*internal_models.TicketStatus, group_id uint64) []*internal_models.TicketStatus {
	filteredList := make([]*internal_models.TicketStatus, 0)
	if u.permUC.CheckPermission(group_id, global_const.AdminTA) {
		return list
	}

	if u.permUC.CheckPermission(group_id, global_const.TicketTA_Work) {
		for _, stat := range list {
			if stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSResolve].ID ||
				stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSWait].ID {
				continue
			} else {
				filteredList = append(filteredList, stat)
			}
		}
	} else if u.permUC.CheckPermission(group_id, global_const.TicketTA_Resolve) {
		for _, stat := range list {
			if stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSRevision].ID ||
				stat.ID == internal_models.TicketStatusMap[internal_models.KeyTSRejected].ID {
				filteredList = append(filteredList, stat)
			}
		}
	}

	return filteredList
}

func (u *TicketUsecase) CreateTicket(ticket *internal_models.Ticket) (uint64, models.Err) {
	var (
		hash, ticketHash string
		err              models.Err
	)

	if ticket.CatSect, err = u.catSecUC.GetCategorySectionByID(ticket.CatSect.ID); err != nil {
		return 0, err
	}

	ticketHash = ticket.HashСalculation()
	if err := u.store.Get(ticket.Author.Email, &hash); err == nil && hash == ticketHash {
		return 0, nil
	}

	if ticket.CatSect.NeedApproval {
		ticket.Status.Set(internal_models.KeyTSResolve)
	} else {
		ticket.Status.Set(internal_models.KeyTSWait)
	}

	ticket.Date = time.Now().Truncate(time.Millisecond)

	if mask, err := ticket.IpMask(); err != nil {
		ticket.Filial = "not found"
	} else {
		if fil, reg, err := u.regFilUC.GetFilialByIp(mask); err != nil {
			ticket.Filial = "not found"
		} else {
			ticket.Filial = fmt.Sprintf("%s / %s", fil.Name, reg.Name)
		}
	}

	id, er := u.repo.CreateTicket(ticket)
	if er != nil {
		return 0, models.InternalError(er.Error())
	}

	u.store.Add(ticket.Author.Email, ticketHash, viper.GetInt64("app.ttl_cache.life_time"))

	return id, nil
}
