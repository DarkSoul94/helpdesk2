package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
)

var (
	errPermissions_UserUpdate = models.Forbidden("Недостаточно прав для обновления пользователя")

	errPermissions_CreateGroup  = models.Forbidden("Недостаточно прав для создания группы")
	errPermissions_UpdateGroup  = models.Forbidden("Недостаточно прав для обновления группы")
	errPermissions_GetGroupList = models.Forbidden("Недостаточно прав для получения списка групп")
)
