package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
)

type CatSecUsecase struct {
	repo cat_sec_manager.ICatSecRepo
}

func NewCatSecUsecase(repo cat_sec_manager.ICatSecRepo) *CatSecUsecase {
	return &CatSecUsecase{
		repo: repo,
	}
}

func (u *CatSecUsecase) CreateCategory(cat cat_sec_manager.Category) (uint64, models.Err) {
	return 0, nil
}
