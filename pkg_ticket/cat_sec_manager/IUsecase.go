package cat_sec_manager

import "github.com/DarkSoul94/helpdesk2/models"

type ICatSecUsecase interface {
	CreateCategory(cat Category) (uint64, models.Err)
}
