package cat_sec_manager

import "github.com/shopspring/decimal"

type Category struct {
	CategoryID          uint64
	CategoryName        string
	SignificantCategory bool
	OldCategory         bool
	Price               decimal.Decimal
}
