package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func NewCatSecRepo(db *sql.DB) *CatSecRepo {
	return &CatSecRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *CatSecRepo) Close() error {
	r.db.Close()
	return nil
}
