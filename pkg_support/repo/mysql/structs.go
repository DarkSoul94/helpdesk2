package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Repo struct {
	db *sqlx.DB
}

type dbSupport struct {
	SupportID uint64 `db:"support_id"`
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
	InternalNumber string          `db:"internal_number"`
	MobileNumber   string          `db:"mobile_number"`
	BirthDate      string          `db:"birth_date"`
	IsSenior       bool            `db:"is_senior"`
	SeniorID       sql.NullInt64   `db:"senior_id"`
	Wager          decimal.Decimal `db:"wager"`
	Comment        string          `db:"comment"`
	Color          string          `db:"color"`
}
