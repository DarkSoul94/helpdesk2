package mysql

import "github.com/DarkSoul94/helpdesk2/pkg_reports/internal_models"

func (r *ReportsRepo) toModelsReturnedTicket(ticket dbReturnedTicket) internal_models.ReturnedTicket {
	mTicekt := internal_models.ReturnedTicket{
		TicketID:   ticket.TicketID,
		TicketDate: ticket.TicketDate,
		Category:   ticket.Category,
		Section:    ticket.Section,
		TicketText: ticket.TicketText,
		Status:     ticket.TicketText,
		Author:     ticket.Author,
		Support:    ticket.Support,
	}

	if ticket.TicketGrade.Valid {
		mTicekt.TicketGrade = uint64(ticket.TicketGrade.Int64)
	}

	return mTicekt
}
