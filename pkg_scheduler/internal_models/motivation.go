package internal_models

import "github.com/shopspring/decimal"

type Motivation struct {
	SupportID   uint64
	SupportName string
	Color       string
	Motivation  decimal.Decimal
}
