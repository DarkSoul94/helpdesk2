package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"
)

type dbConst struct {
	Name      string         `db:"name"`
	Data      string         `db:"data"`
	DataType  string         `db:"data_type"`
	TableName sql.NullString `db:"table_name"`
}

//dbConstHistory ...
type dbConstHistory struct {
	ID      uint64 `db:"id"`
	Date    string `db:"date"`
	Name    string `db:"name"`
	Val     string `db:"val"`
	ValType string `db:"val_type"`
}

func createConstHistory(name string, data interface{}) dbConstHistory {
	return dbConstHistory{
		Name:    name,
		Date:    time.Now().Format("2006-01") + "-01",
		Val:     fmt.Sprintf("%v", data),
		ValType: fmt.Sprintf("%T", data),
	}
}

func (c *dbConst) ToConst(name string, data interface{}) {
	*c = dbConst{
		Name:     name,
		Data:     fmt.Sprintf("%v", data),
		DataType: fmt.Sprintf("%T", data),
		TableName: sql.NullString{
			Valid: false,
		},
	}
}

func (c *dbConst) FromConst(target interface{}) error {
	if target != nil {
		v := reflect.ValueOf(target)

		if fmt.Sprintf("%v", v.Elem().Type()) != c.DataType {
			return errors.New("invalid destination type")
		}
		v.Elem().Set(reflect.ValueOf(c.Data))
	}
	return nil
}
