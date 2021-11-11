package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

var (
	AuthErr_CreateToken = models.InternalError("Failed create token for user")
	AuthErr_Login       = models.Unauthorized("Wrong e-mail or password")
)
