package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/reg_fil_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_user"
)

type TicketUsecase struct {
	repo     pkg_ticket.ITicketRepo
	catSecUC cat_sec_manager.ICatSecUsecase
	regFilUC reg_fil_manager.IRegFilUsecase
	userUC   pkg_user.IUserUsecase
}

func NewTicketUsecase(
	tRepo pkg_ticket.ITicketRepo,
	catSecUC cat_sec_manager.ICatSecUsecase,
	regFilUC reg_fil_manager.IRegFilUsecase,
	userUC pkg_user.IUserUsecase,
) *TicketUsecase {
	return &TicketUsecase{
		repo:     tRepo,
		catSecUC: catSecUC,
		regFilUC: regFilUC,
		userUC:   userUC,
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

	}

	return list, nil
}
