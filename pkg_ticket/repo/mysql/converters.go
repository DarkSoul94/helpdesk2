package mysql

import (
	"database/sql"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

func (r *TicketRepo) toModelTicketStatus(stat dbTicketStatus) *internal_models.TicketStatus {
	return &internal_models.TicketStatus{
		ID:   stat.ID,
		Name: stat.Name,
	}
}

func (r *TicketRepo) toDbTicketStatus(stat *internal_models.TicketStatus) *dbTicketStatus {
	return &dbTicketStatus{
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
		Status:    r.toDbTicketStatus(ticket.Status),
		Grade:     sql.NullInt32{Int32: int32(ticket.Grade), Valid: true},
	}

	if len(ticket.Filial) > 0 {
		dbTick.Filial.String = ticket.Filial
		dbTick.Filial.Valid = true
	} else {
		dbTick.Filial.Valid = false
	}

	if len(ticket.IP) > 0 {
		dbTick.IP.String = ticket.IP
		dbTick.IP.Valid = true
	} else {
		dbTick.IP.Valid = false
	}

	if ticket.Author != nil && ticket.Author.ID != 0 {
		dbTick.AuthorID.Int64 = int64(ticket.Author.ID)
		dbTick.AuthorID.Valid = true
	} else {
		dbTick.AuthorID.Valid = false
	}

	if ticket.Support != nil && ticket.Support.ID != 0 {
		dbTick.SupportID.Int64 = int64(ticket.Support.ID)
		dbTick.SupportID.Valid = true
	} else {
		dbTick.SupportID.Valid = false
	}

	if ticket.ResolvedUser != nil && ticket.ResolvedUser.ID != 0 {
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

func (r *TicketRepo) toModelTicket(dbTick dbTicket) *internal_models.Ticket {
	mTick := &internal_models.Ticket{
		ID:      dbTick.ID,
		Date:    dbTick.Date,
		Text:    dbTick.Text,
		CatSect: &internal_models.SectionWithCategory{ID: dbTick.SectionID},
		Status:  r.toModelTicketStatus(*dbTick.Status),
	}

	if dbTick.Filial.Valid {
		mTick.Filial = dbTick.Filial.String
	}

	if dbTick.IP.Valid {
		mTick.IP = dbTick.IP.String
	}

	if dbTick.AuthorID.Valid {
		mTick.Author = &models.User{ID: uint64(dbTick.AuthorID.Int64)}
	}

	if dbTick.SupportID.Valid {
		mTick.Support = &models.User{ID: uint64(dbTick.SupportID.Int64)}
	}

	if dbTick.ResolvedUserID.Valid {
		mTick.ResolvedUser = &models.User{ID: uint64(dbTick.ResolvedUserID.Int64)}
	}

	if dbTick.ServiceComment.Valid {
		mTick.ServiceComment = dbTick.ServiceComment.String
	}

	if dbTick.Grade.Valid {
		mTick.Grade = uint(dbTick.Grade.Int32)
	}

	return mTick
}

func (r *TicketRepo) toModelTicketStatusHistory(history dbTicketStatusHistory) *internal_models.TicketStatusHistory {
	return &internal_models.TicketStatusHistory{
		ID:          history.ID,
		TicketId:    history.TicketId,
		ChangedUser: &models.User{ID: history.ChangedUserID},
		SelectTime:  history.SelectTime,
		Status:      r.toModelTicketStatus(*history.Status),
		Duration:    history.Duration,
	}
}

func (r *TicketRepo) toDbTicketStatusHistory(history *internal_models.TicketStatusHistory) dbTicketStatusHistory {
	return dbTicketStatusHistory{
		ID:            history.ID,
		TicketId:      history.TicketId,
		ChangedUserID: history.ChangedUser.ID,
		SelectTime:    history.SelectTime,
		Status:        r.toDbTicketStatus(history.Status),
		Duration:      history.Duration,
	}
}
