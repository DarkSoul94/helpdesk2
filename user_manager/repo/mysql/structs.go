package mysql

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
}

type dbGroup struct {
	ID          uint64 `db:"id"`
	Name        string `db:"name"`
	Permissions []byte `db:"permissions"`
}

type dbUser struct {
	ID         uint64 `db:"id"`
	Email      string `db:"email"`
	Name       string `db:"name"`
	GroupID    uint64 `db:"group_id"`
	Department string `db:"department"`
}
