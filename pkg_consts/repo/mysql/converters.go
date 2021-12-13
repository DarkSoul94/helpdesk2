package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

var (
	decimalType = reflect.TypeOf(decimal.Zero)
	float64Type = reflect.TypeOf(float64(0))
	uint64Type  = reflect.TypeOf(uint64(0))
	stringType  = reflect.TypeOf("")
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

		switch v.Elem().Type() {
		case stringType:
			v.Elem().Set(reflect.ValueOf(c.Data))

		case float64Type:
			temp, _ := strconv.ParseFloat(c.Data, 64)
			v.Elem().Set(reflect.ValueOf(temp))

		case uint64Type:
			temp, _ := strconv.ParseUint(c.Data, 10, 64)
			v.Elem().Set(reflect.ValueOf(temp))

		case decimalType:
			temp, _ := decimal.NewFromString(c.Data)
			v.Elem().Set(reflect.ValueOf(temp))

		default:
			return errors.New("Тип контстанты не определен в системе")
		}
	}
	return nil
}
