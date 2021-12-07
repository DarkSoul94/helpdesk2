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

type dbReturnedTicket struct {
	TicketID    uint64        `db:"ticket_id"`
	TicketDate  time.Time     `db:"ticket_date"`
	Category    string        `db:"category"`
	Section     string        `db:"section"`
	TicketText  string        `db:"ticket_text"`
	Status      string        `db:"status"`
	Author      string        `db:"author"`
	Support     string        `db:"support"`
	TicketGrade sql.NullInt64 `db:"ticket_grade"`
}

type dbTicketCount struct {
	Day   time.Time `db:"day"`
	Hour  time.Time `db:"hour"`
	Count uint      `db:"count"`
}

type dbSupportStatusesHistoryPerWeekDays struct {
	WeekDay     uint          `db:"week_day"`
	StatusName  string        `db:"status_name"`
	SupportName string        `db:"support_name"`
	Duration    time.Duration `db:"duration"`
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
