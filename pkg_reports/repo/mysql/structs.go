package mysql

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ReportsRepo struct {
	db *sqlx.DB
}

type dbAverageGrade struct {
	SupportName  string  `db:"user_name"`
	AverageGrade float64 `db:"average_grade"`
}

type dbSupportStatusHistory struct {
	SupportName string        `db:"support_name"`
	StatusName  string        `db:"status_name"`
	SelectTime  time.Time     `db:"select_time"`
	Duration    time.Duration `db:"duration"`
}
