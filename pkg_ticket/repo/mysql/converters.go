package mysql

import (
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

func (r *TicketRepo) toModelTicketStatus(stat dbTicketStatus) *internal_models.TicketStatus {
	return &internal_models.TicketStatus{
		ID:   stat.ID,
		Name: stat.Name,
	}
}
