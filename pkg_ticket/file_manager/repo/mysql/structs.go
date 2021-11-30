package mysql

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type FileRepo struct {
	db *sqlx.DB
}

type dbFile struct {
	ID        uint64         `db:"file_id"`
	Name      string         `db:"file_name"`
	Date      time.Time      `db:"file_date"`
	Data      sql.NullString `db:"file_data"`
	Extension string         `db:"file_extension"`
	Path      sql.NullString `db:"path"`
	TicketId  uint64         `db:"ticket_id"`
}
