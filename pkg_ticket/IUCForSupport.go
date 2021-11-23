package pkg_ticket

import "github.com/DarkSoul94/helpdesk2/models"

type IUCForSupport interface {
	ResetSupportInTickets(ticketIDs []uint64) models.Err
	GetTicketsCounts(supportID uint64, keys ...string) map[string]int
	GetTodayTicketsCounts(supportID uint64, keys ...string) map[string]int
}
