package mysql

import "github.com/DarkSoul94/helpdesk2/models"

var (
	errSupportCreate = models.InternalError("Не удалось создать нового саппорта")
	errSupportDelete = models.InternalError("Не удалось удалить пользователя из списка саппортов")
	errSupportGet    = models.BadRequest("Суппорт с таким ID не найден в базе")

	errCardCreate = models.InternalError("Не удалось создать карточку саппорта")
	errCardDelete = models.InternalError("Не удалось удалить карточку саппорта")
	errCardGet    = models.InternalError("Не удалось получить карточку саппорта")
)
