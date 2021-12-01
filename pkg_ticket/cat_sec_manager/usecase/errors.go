package usecase

import "github.com/DarkSoul94/helpdesk2/models"

var (
	ErrCategoryExist    = models.BadRequest("Категория уже существует")
	ErrCategoryNotExist = models.BadRequest("Указанной категории не существует")

	ErrCategorySectionExist    = models.BadRequest("Раздел категории уже существует")
	ErrCategorySectionNotExist = models.BadRequest("Указанного раздела категории не существует")
)
