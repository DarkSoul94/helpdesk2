package mysql

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type ReportsRepo struct {
	db *sqlx.DB
}

type dbTicketStatusDifference struct {
	TicketID    uint64        `db:"ticket_id"`
	SupportName string        `db:"support_name"`
	Section     string        `db:"section"`
	StatusName  string        `db:"status_name"`
	Duration    time.Duration `db:"duration"`
}

type dbAverageGrade struct {
	SupportName  string  `db:"user_name"`
	AverageGrade float64 `db:"average_grade"`
}

type dbTicketsGrade struct {
	TicketID   uint64 `db:"ticket_id"`
	Grade      uint   `db:"ticket_grade"`
	UserName   string `db:"user_name"`
	Department string `db:"department"`
}

type dbTicketCount struct {
	Day   time.Time `db:"day"`
	Hour  time.Time `db:"hour"`
	Count uint      `db:"count"`
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
