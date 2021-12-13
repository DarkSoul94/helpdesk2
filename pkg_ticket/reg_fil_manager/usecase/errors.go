package usecase

import "github.com/DarkSoul94/helpdesk2/models"

var (
	ErrRegionExist    = models.BadRequest("Регион уже существует")
	ErrRegionNotExist = models.BadRequest("Указанного региона не существует")

	ErrFilialExist    = models.BadRequest("Филиал уже существует")
	ErrFilialNotExist = models.BadRequest("Указанного филиала не существует")
)
