package internal_models

import "github.com/shopspring/decimal"

type Category struct {
	ID          uint64
	Name        string
	Significant bool
	Old         bool
	Price       decimal.Decimal
}
