package mysql

import "github.com/jmoiron/sqlx"

type ReportsRepo struct {
	db *sqlx.DB
}

type dbAverageGrade struct {
	SupportName  string  `db:"user_name"`
	AverageGrade float64 `db:"average_grade"`
}
