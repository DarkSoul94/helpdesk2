package usecase

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type Usecase struct {
	repo pkg_ticket.IRepoForSupport
}

func NewTicketUCForSupport(repo pkg_ticket.IRepoForSupport) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ResetSupportInTickets(ticketIDs []uint64) models.Err {
	//TODO написать функцию
	return nil
}

//GetTicketsCounts возвращает соответствие ключей статусов запросов и количества этих запросов за все время по саппорту
func (u *Usecase) GetTicketsCounts(supportID uint64, keys ...string) map[string]int {
	result := make(map[string]int)
	for _, key := range keys {
		result[key] = u.repo.GetTicketsCount(supportID, internal_models.TicketStatusMap[key].ID)
	}
	return result
}

//GetTodayTicketsCounts возвращает соответствие ключей статусов запросов и количества этих запросов за сегодня по саппорту
func (u *Usecase) GetTodayTicketsCounts(supportID uint64, keys ...string) map[string]int {
	result := make(map[string]int)
	for _, key := range keys {
		result[key] = u.repo.GetTodayTicketsCount(supportID, internal_models.TicketStatusMap[key].ID)
	}
	return result
}
