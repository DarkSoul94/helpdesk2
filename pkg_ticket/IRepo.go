package pkg_ticket

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type ITicketRepo interface {
	GetTicketStatuses() ([]*internal_models.TicketStatus, error)
	CreateTicket(ticket *internal_models.Ticket) (uint64, error)

	Close() error
}
