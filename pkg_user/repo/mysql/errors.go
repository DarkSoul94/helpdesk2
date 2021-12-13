package mysql

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

var (
	userErr_Create  = models.InternalError("Не удалось записать нового пользователя")
	userErr_Update  = models.InternalError("Не удалось обновить пользователя")
	userErr_Get     = models.InternalError("Не удалось получить пользователя")
	userErr_GetList = models.InternalError("Не удалось получить список пользователей")

	commonErr_Read = models.InternalError("Не удалось получить данные")
)
