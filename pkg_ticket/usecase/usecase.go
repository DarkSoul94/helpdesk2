package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type TicketUsecase struct {
	ticketRepo pkg_ticket.ITicketRepo
	catSecUC   cat_sec_manager.ICatSecUsecase
}

func NewTicketUsecase(
	tRepo pkg_ticket.ITicketRepo,
	catSecUC cat_sec_manager.ICatSecUsecase,
) *TicketUsecase {
	return &TicketUsecase{
		ticketRepo: tRepo,
		catSecUC:   catSecUC,
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
