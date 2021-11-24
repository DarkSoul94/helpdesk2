package pkg_ticket

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type ITicketRepo interface {
	GetTicketStatuses() ([]*internal_models.TicketStatus, error)
	GetTicketStatusesSortPriority(isSupport bool) map[uint]uint
	GetLastTicketStatusHistory(ticketID uint64) (*internal_models.TicketStatusHistory, error)
	CreateTicketStatusHistory(history *internal_models.TicketStatusHistory) error
	UpdateTicketStatusHistory(history *internal_models.TicketStatusHistory) error

	CreateTicket(ticket *internal_models.Ticket) (uint64, error)
	GetTicketListForAdmin(limit, offset int) ([]*internal_models.Ticket, error)
	GetTicketListForSupport(supportID uint64, limit, offset int) ([]*internal_models.Ticket, error)
	GetTicketListForUser(authorID uint64, limit, offset int) ([]*internal_models.Ticket, error)
	GetTicketListForApproval(groupID uint64, limit, offset int, forResolver bool) ([]*internal_models.Ticket, error)
	CheckNeedApprovalTicketExist(groupID uint64, forResolver bool) (bool, error)

	Close() error
}
