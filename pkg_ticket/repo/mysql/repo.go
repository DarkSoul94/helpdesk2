package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func NewTicketRepo(db *sql.DB) *TicketRepo {
	return &TicketRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *TicketRepo) Close() error {
	r.db.Close()
	return nil
}
