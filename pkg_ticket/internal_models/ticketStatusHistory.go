package internal_models

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
)

type TicketStatusHistory struct {
	ID          uint64
	TicketId    uint64
	ChangedUser *models.User
	SelectTime  time.Time
	Status      *TicketStatus
	Duration    time.Duration
}

func (th *TicketStatusHistory) New(ticketID, changedUserID uint64, newStatus *TicketStatus, currentTime time.Time) {
	*th = TicketStatusHistory{
		TicketId:    ticketID,
		ChangedUser: &models.User{ID: changedUserID},
		SelectTime:  currentTime,
		Status:      newStatus,
		Duration:    0,
	}
}
