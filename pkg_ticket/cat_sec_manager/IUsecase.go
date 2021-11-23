package cat_sec_manager

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type ICatSecUsecase interface {
	CreateCategory(cat *internal_models.Category) (uint64, models.Err)
	UpdateCategory(cat *internal_models.Category) models.Err

	CreateCategorySection(sec *internal_models.CategorySection) (uint64, models.Err)
	UpdateCategorySection(sec *internal_models.CategorySection) models.Err
	GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, models.Err)
	GetCategorySectionList() ([]internal_models.CategorySectionList, models.Err)
	GetCategorySectionByID(id uint64) (*internal_models.CategorySection, models.Err)
	GetSectionWithCategoryByID(id uint64) (*internal_models.SectionWithCategory, models.Err)
}
