package mysql

import (
	"database/sql"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

func (r *TicketRepo) toModelTicketStatus(stat dbTicketStatus) *internal_models.TicketStatus {
	return &internal_models.TicketStatus{
		ID:   stat.ID,
		Name: stat.Name,
	}
}

func (r *TicketRepo) toDbTicket(ticket *internal_models.Ticket) dbTicket {
	dbTick := dbTicket{
		ID:        ticket.ID,
		Date:      ticket.Date,
		SectionID: ticket.CatSect.ID,
		Text:      ticket.Text,
		StatusID:  ticket.Status.ID,
		Grade:     sql.NullInt32{Int32: int32(ticket.Grade), Valid: true},
	}

	if len(ticket.Filial) > 0 {
		dbTick.Filial.String = ticket.Filial
		dbTick.Filial.Valid = true
	} else {
		dbTick.Filial.Valid = false
	}

	if len(ticket.Ip) > 0 {
		dbTick.IP.String = ticket.Ip
		dbTick.IP.Valid = true
	} else {
		dbTick.IP.Valid = false
	}

	if ticket.Author != nil {
		dbTick.AuthorID.Int64 = int64(ticket.Author.ID)
		dbTick.AuthorID.Valid = true
	} else {
		dbTick.AuthorID.Valid = false
	}

	if ticket.Support != nil {
		dbTick.SupportID.Int64 = int64(ticket.Support.ID)
		dbTick.SupportID.Valid = true
	} else {
		dbTick.SupportID.Valid = false
	}

	if ticket.ResolvedUser != nil {
		dbTick.ResolvedUserID.Int64 = int64(ticket.ResolvedUser.ID)
		dbTick.ResolvedUserID.Valid = true
	} else {
		dbTick.ResolvedUserID.Valid = false
	}

	if len(ticket.ServiceComment) > 0 {
		dbTick.ServiceComment.String = ticket.ServiceComment
		dbTick.ServiceComment.Valid = true
	} else {
		dbTick.ServiceComment.Valid = false
	}

	return dbTick
}
