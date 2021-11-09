package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

var (
	AuthErr_CreateToken = models.BadRequest("Failed create token for user")
	AuthErr_Login       = models.BadRequest("Wrong e-mail or password")
)
