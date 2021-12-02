package mysql

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Repo struct {
	db *sqlx.DB
}

type dbSupport struct {
	SupportID uint64 `db:"support_id"`
	Name      string `db:"support_name"`
	StatusID  uint64 `db:"status_id"`
	Priority  bool   `db:"priority"`
}

type dbStatus struct {
	ID           uint64 `db:"support_status_id"`
	Name         string `db:"support_status_name"`
	AcceptTicket bool   `db:"accept_ticket"`
}

type dbCard struct {
	ID             uint64          `db:"id"`
	SupportID      uint64          `db:"support_id"`
	SupportName    string          `db:"support_name"`
	InternalNumber string          `db:"internal_number"`
	MobileNumber   string          `db:"mobile_number"`
	BirthDate      string          `db:"birth_date"`
	IsSenior       bool            `db:"is_senior"`
	SeniorID       sql.NullInt64   `db:"senior_id"`
	SeniorName     sql.NullString  `db:"senior_name"`
	Wager          decimal.Decimal `db:"wager"`
	Comment        string          `db:"comment"`
	Color          string          `db:"color"`
}

type dbShift struct {
	ID            uint64       `db:"id"`
	SupportID     uint64       `db:"support_id"`
	OpeningTime   time.Time    `db:"opening_time"`
	ClosingTime   sql.NullTime `db:"closing_time"`
	ClosingStatus bool         `db:"closing_status"`
}

type dbStatusHistory struct {
	ID         uint64        `db:"id"`
	SupportID  uint64        `db:"support_id"`
	StatusID   uint64        `db:"status_id"`
	SelectTime time.Time     `db:"select_time"`
	Duration   time.Duration `db:"duration"`
	ShiftID    uint64        `db:"shift_id"`
}
