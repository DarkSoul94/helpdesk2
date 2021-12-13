package mysql

import "database/sql"

type dBConst struct {
	Name      string         `db:"name"`
	Data      string         `db:"data"`
	DataType  string         `db:"data_type"`
	TableName sql.NullString `db:"table_name"`
}

type dbHistory struct {
	ID      uint64 `db:"id"`
	Date    string `db:"date"`
	Name    string `db:"name"`
	Val     string `db:"val"`
	ValType string `db:"val_type"`
}