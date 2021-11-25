package usecase

import "github.com/DarkSoul94/helpdesk2/models"

var (
	ErrConnotUpdateTicket = models.Forbidden("В данный момент вы не можете работать с этим запросом")
)
