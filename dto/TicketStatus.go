package dto

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

type OutTicketStatus struct {
	ID   uint64 `json:"ticket_status_id"`
	Name string `json:"ticket_status_name"`
}

func ToOutTicketStatus(stat *internal_models.TicketStatus) OutTicketStatus {
	return OutTicketStatus{
		ID:   stat.ID,
		Name: stat.Name,
	}
}
