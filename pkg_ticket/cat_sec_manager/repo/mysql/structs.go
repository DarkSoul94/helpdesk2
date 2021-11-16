package mysql

import "github.com/jmoiron/sqlx"

type CatSecRepo struct {
	db *sqlx.DB
}
