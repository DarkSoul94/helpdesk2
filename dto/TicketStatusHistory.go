package dto

import (
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type OutTicketStatusHistory struct {
	SelectTime  time.Time `json:"curr_status_time"`
	Status      string    `json:"curr_status"`
	ChangedUser string    `json:"changed_user"`
	Duration    uint64    `json:"difference"`
}

func ToOutTicketStatusHistory(history *internal_models.TicketStatusHistory) OutTicketStatusHistory {
	outHistory := OutTicketStatusHistory{
		SelectTime:  history.SelectTime,
		Status:      history.Status.Name,
		ChangedUser: history.ChangedUser.Name,
	}

	if history.Duration == 0 {
		outHistory.Duration = uint64(time.Now().Local().Sub(history.SelectTime).Seconds())
	} else {
		outHistory.Duration = uint64(history.Duration.Seconds())
	}

	return outHistory
}
