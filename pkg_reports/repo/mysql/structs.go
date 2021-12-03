package mysql

import (
	"database/sql"
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

type dbSupportsShifts struct {
	Support            string        `db:"support"`
	OpeningDate        time.Time     `db:"opening_time"`
	ClosingDate        sql.NullTime  `db:"closing_time"`
	CountOfMinutesLate sql.NullInt64 `db:"count_of_minutes_late"`
}

type dbSupportStatusHistory struct {
	SupportName string        `db:"support_name"`
	StatusName  string        `db:"status_name"`
	SelectTime  time.Time     `db:"select_time"`
	Duration    time.Duration `db:"duration"`
}
