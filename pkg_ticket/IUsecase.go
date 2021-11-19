package pkg_ticket

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type ITicketUsecase interface {
	CreateCategory(cat *internal_models.Category) (uint64, models.Err)
	UpdateCategory(cat *internal_models.Category) models.Err

	CreateCategorySection(sec *internal_models.CategorySection) (uint64, models.Err)
	UpdateCategorySection(sec *internal_models.CategorySection) models.Err
	GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, models.Err)
	GetCategorySectionList() ([]internal_models.CategorySectionList, models.Err)

	CreateRegion(reg *internal_models.Region) (uint64, models.Err)
	UpdateRegion(reg *internal_models.Region) models.Err
	DeleteRegion(id uint64) models.Err

	CreateFilial(fil *internal_models.Filial) (uint64, models.Err)
	UpdateFilial(fil *internal_models.Filial) models.Err
	DeleteFilial(id uint64) models.Err
	GetRegionsWithFilials() ([]*internal_models.RegionWithFilials, models.Err)

	GetTicketStatuses(group_id uint64, all bool) ([]*internal_models.TicketStatus, models.Err)
}