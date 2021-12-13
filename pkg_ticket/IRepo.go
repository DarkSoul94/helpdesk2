package pkg_ticket

import (
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type ITicketRepo interface {
	GetTicketStatuses() ([]*internal_models.TicketStatus, error)
	GetTicketStatusesSortPriority(isSupport bool) map[uint]uint
	GetLastTicketStatusHistory(ticketID uint64) (*internal_models.TicketStatusHistory, error)
	CreateTicketStatusHistory(history *internal_models.TicketStatusHistory) error
	UpdateTicketStatusHistory(history *internal_models.TicketStatusHistory) error
	GetAllTicketStatusHistory(ticketID uint64) ([]*internal_models.TicketStatusHistory, error)

	CreateTicket(ticket *internal_models.Ticket) (uint64, error)
	UpdateTicket(ticket *internal_models.Ticket) error
	GetTicketsCount(supportID, statusID uint64) int
	GetTicketListForAdmin(limit, offset int) ([]*internal_models.Ticket, error)
	GetTicketListForSupport(supportID uint64, limit, offset int) ([]*internal_models.Ticket, error)
	GetTicketListForUser(authorID uint64, limit, offset int) ([]*internal_models.Ticket, error)
	GetTicketListForApproval(groupID uint64, limit, offset int, forResolver bool) ([]*internal_models.Ticket, error)
	GetTicketListForDistribute() ([]*internal_models.Ticket, error)
	GetTicketListForReturnToDistribute() ([]*internal_models.Ticket, error)
	FilterDispatcher(filter map[string]interface{}) (string, []interface{})
	GetFilteredTicketsList(query string, args []interface{}, fullSearch bool) ([]*internal_models.Ticket, error)
	GetTicket(ticketID uint64) (*internal_models.Ticket, error)
	StealTicket(ticketID, supportID uint64, toWork bool) error
	TicketGrade(ticketID, userID uint64, grade uint) error
	CheckNeedApprovalTicketExist(groupID uint64, forResolver bool) (bool, error)

	Close() error
}
