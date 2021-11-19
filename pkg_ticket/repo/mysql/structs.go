package mysql

import "github.com/jmoiron/sqlx"

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
