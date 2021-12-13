package http

import "github.com/DarkSoul94/helpdesk2/models"

var (
	errBlankField = models.BadRequest("Одно или несколько обязательных полей - пустое")
)
