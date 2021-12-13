package mysql

import (
	"database/sql"
	"time"
)

type dbOffice struct {
	ID      uint64 `db:"id"`
	Name    string `db:"name"`
	Color   string `db:"color"`
	Deleted bool   `db:"deleted"`
}

type dbCell struct {
	ID        uint64        `db:"id"`
	SupportID uint64        `db:"support_id"`
	OfficeID  sql.NullInt64 `db:"office_id"`
	StartTime string        `db:"start_time"`
	EndTime   string        `db:"end_time"`
	Date      time.Time     `db:"date"`
	Vacation  bool          `db:"vacation"`
	SickLeave bool          `db:"sick_leave"`
}

type dbLateness struct {
	ID          uint64       `db:"id"`
	Date        time.Time    `db:"date"`
	SupportID   uint64       `db:"support_id"`
	SupportName string       `db:"support_name"`
	Cause       string       `db:"cause"`
	Decision    sql.NullBool `db:"decision"`
	Difference  uint64       `db:"difference"`
}
