package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func (c *dBConst) ToConst(name string, data interface{}) {
	*c = dBConst{
		Name:     name,
		Data:     fmt.Sprintf("%v", data),
		DataType: fmt.Sprintf("%T", data),
		TableName: sql.NullString{
			Valid: false,
		},
	}
}

func (c *dbHistory) ToHistory(name string, data interface{}) {
	*c = dbHistory{
		Date:    time.Now().Format("2006-01") + "-01",
		Name:    name,
		Val:     fmt.Sprintf("%v", data),
		ValType: fmt.Sprintf("%T", data),
	}
}

func (c *dBConst) FromConst(target interface{}) error {
	if target != nil {
		v := reflect.ValueOf(target)

		if fmt.Sprintf("%v", v.Elem().Type()) != c.DataType {
			return errors.New("invalid destination type")
		}
		v.Elem().Set(reflect.ValueOf(c.Data))
	}
	return nil
}
