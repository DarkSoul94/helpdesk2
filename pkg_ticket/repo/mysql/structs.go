package mysql

import "github.com/jmoiron/sqlx"

type TicketRepo struct {
	db *sqlx.DB
}
