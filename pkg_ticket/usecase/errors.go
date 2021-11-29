package usecase

import "github.com/DarkSoul94/helpdesk2/models"

var (
	errCannotUpdateTicket   = models.Forbidden("В данный момент вы не можете работать с этим запросом")
	errTicketOld            = models.InternalError("С этим запрососм нельзя работать")
	errCannotSelectStatus   = models.Forbidden("Вы не можете выбрать этот статус запроса")
	errTicketWithoutSupport = models.BadRequest("За запросом не закреплен сотрудник ТП")
	errLimitImplementation  = models.BadRequest("Превышен лимит по количеству запрос в статусе 'В процесе реализации' у сотрудника ТП")
	errLimitPostponed       = models.BadRequest("Превышен лимит по количеству запрос в статусе 'Отложен' у сотрудника ТП")
)
