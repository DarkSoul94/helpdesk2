package mysql

import "database/sql"

type dBConst struct {
	Name      string         `db:"name"`
	Data      string         `db:"data"`
	DataType  string         `db:"data_type"`
	TableName sql.NullString `db:"table_name"`
}
