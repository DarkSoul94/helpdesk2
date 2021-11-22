package usecase

import "github.com/DarkSoul94/helpdesk2/models"

var (
	supportErr_ClosedShift  = models.BadRequest("Смена саппорта закрыта")
	supportErr_AlreadyOpen  = models.BadRequest("Смена саппорта уже открыта")
	supportErr_CannotReopen = models.BadRequest("Вы не можете открыть смену указанному саппорту")
	supportErr_Busy         = models.BadRequest("У саппорта есть запросы в работе")
)
