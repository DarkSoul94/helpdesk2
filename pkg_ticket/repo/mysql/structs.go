package mysql

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type TicketRepo struct {
	db *sqlx.DB
}

type dbTicketStatus struct {
	ID               uint64 `db:"ticket_status_id"`
	Name             string `db:"ticket_status_name"`
	NotDisplay       bool   `db:"not_display"`
	SortPrioritySupp uint   `db:"sort_priority_supp"`
	SortPriorityUser uint   `db:"sort_priority_user"`
}

type dbTicket struct {
	ID             uint64          `db:"ticket_id"`
	Date           time.Time       `db:"ticket_date"`
	SectionID      uint64          `db:"section_id"`
	Text           string          `db:"ticket_text"`
	Status         *dbTicketStatus `db:""`
	Filial         sql.NullString  `db:"filial"`
	IP             sql.NullString  `db:"ip"`
	AuthorID       sql.NullInt64   `db:"ticket_author_id"`
	SupportID      sql.NullInt64   `db:"support_id"`
	ResolvedUserID sql.NullInt64   `db:"resolved_user_id"`
	ServiceComment sql.NullString  `db:"service_comment"`
	Grade          sql.NullInt32   `db:"ticket_grade"`
}
