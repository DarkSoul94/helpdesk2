package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type CatSecUsecase struct {
	repo cat_sec_manager.ICatSecRepo
}

func NewCatSecUsecase(repo cat_sec_manager.ICatSecRepo) *CatSecUsecase {
	return &CatSecUsecase{
		repo: repo,
	}
}

func (u *CatSecUsecase) CreateCategory(cat *internal_models.Category) (uint64, models.Err) {
	if u.repo.CheckCategoryExist(0, cat.Name) {
		return 0, ErrCategoryExist
	}

	id, err := u.repo.CreateCategory(cat)
	if err != nil {
		return 0, models.InternalError(err.Error())
	}

	return id, nil
}

func (u *CatSecUsecase) UpdateCategory(cat *internal_models.Category) models.Err {
	if !u.repo.CheckCategoryExist(cat.ID, "") {
		return ErrCategoryNotExist
	}

	err := u.repo.UpdateCategory(cat)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *CatSecUsecase) CreateCategorySection(sec *internal_models.CategorySection) (uint64, models.Err) {
	if u.repo.CheckCategorySectionExist(0, sec.CategoryID, sec.Name) {
		return 0, ErrCategorySectionExist
	}

	id, err := u.repo.CreateCategorySection(sec)
	if err != nil {
		return 0, models.InternalError(err.Error())
	}

	return id, nil
}

func (u *CatSecUsecase) UpdateCategorySection(sec *internal_models.CategorySection) models.Err {
	if !u.repo.CheckCategorySectionExist(sec.ID, 0, "") {
		return ErrCategorySectionExist
	}

	err := u.repo.UpdateCategorySection(sec)
	if err != nil {
		return models.InternalError(err.Error())
	}

	return nil
}

func (u *CatSecUsecase) GetCategorySection(forSearch bool) ([]*internal_models.SectionWithCategory, models.Err) {
	list, err := u.repo.GetCategorySection(forSearch)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return list, nil
}

func (u *CatSecUsecase) GetCategorySectionList() ([]internal_models.CategorySectionList, models.Err) {
	list, err := u.repo.GetCategorySectionList()
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return list, nil
}

func (u *CatSecUsecase) GetCategorySectionByID(id uint64) (*internal_models.CategorySection, models.Err) {
	if !u.repo.CheckCategorySectionExist(id, 0, "") {
		return nil, ErrCategorySectionExist
	}

	sect, err := u.repo.GetCategorySectionByID(id)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return sect, nil
}

func (u *CatSecUsecase) GetSectionWithCategoryByID(id uint64) (*internal_models.SectionWithCategory, models.Err) {
	if !u.repo.CheckCategorySectionExist(id, 0, "") {
		return nil, ErrCategorySectionNotExist
	}

	sect, err := u.repo.GetSectionWithCategoryByID(id)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return sect, nil
}
