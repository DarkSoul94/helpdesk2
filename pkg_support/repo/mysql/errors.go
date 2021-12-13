package mysql

import "github.com/DarkSoul94/helpdesk2/models"

var (
	errSupportCreate         = models.InternalError("Не удалось создать нового саппорта")
	errSupportUpdate         = models.InternalError("Не удалось обновить статус саппорта")
	errSupportDelete         = models.InternalError("Не удалось удалить пользователя из списка саппортов")
	errSupportGetRandom      = models.InternalError("Не удалось получить случайного свободного саппорта")
	errSupportModifyActivity = models.InternalError("Не удалось модифицировать запись о распределении запроса на саппорта")
	errSupportGetList        = models.InternalError("Не удалось получить список саппортов")
	errSupportGetActive      = models.BadRequest("Не найдены активные саппорты")
	errSupportGet            = models.BadRequest("Суппорт с таким ID не найден в базе")

	errShiftGet         = models.InternalError("Не удалось получить данные по последней смене саппорта")
	errShiftUpdateShift = models.InternalError("Не удалось обновить смену саппорту")

	errStatusGet = models.InternalError("Не удалось получить статус саппорта")

	errHistoryGet    = models.InternalError("Не удалось получить запись из истории статусов саппортов")
	errHistoryCreate = models.InternalError("Не удалось создать запись в истории статусов саппортов")

	errCardCreate = models.InternalError("Не удалось создать карточку саппорта")
	errCardUpdate = models.InternalError("Не удалось обновить карточку саппорта")
	errCardDelete = models.InternalError("Не удалось удалить карточку саппорта")
	errCardGet    = models.InternalError("Не удалось получить карточку саппорта")
)
