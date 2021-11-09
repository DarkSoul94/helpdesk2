package mysql

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

var (
	UserErr_Create  = models.InternalError("Failed insert new user to db")
	UserErr_Update  = models.InternalError("Failed update user")
	UserErr_Get     = models.InternalError("Failed read user from db")
	UserErr_GetList = models.InternalError("Failed read users list from db")
)
