package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

type dbUser struct {
	ID         uint64         `db:"user_id"`
	Name       string         `db:"user_name"`
	Email      string         `db:"email"`
	GroupID    uint64         `db:"group_id"`
	Department sql.NullString `db:"department"`
}
