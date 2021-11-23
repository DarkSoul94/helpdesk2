package cat_sec_manager

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type ICatSecRepo interface {
	CheckCategoryExist(id uint64, name string) bool
	CreateCategory(cat *internal_models.Category) (uint64, error)
	UpdateCategory(cat *internal_models.Category) error

	CheckCategorySectionExist(id, cat_id uint64, name string) bool
	CreateCategorySection(sec *internal_models.CategorySection) (uint64, error)
	UpdateCategorySection(sec *internal_models.CategorySection) error
	GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, error)
	GetCategorySectionList() ([]internal_models.CategorySectionList, error)
	GetCategorySectionByID(id uint64) (*internal_models.CategorySection, error)
	GetSectionWithCategoryByID(id uint64) (*internal_models.SectionWithCategory, error)

	Close() error
}
