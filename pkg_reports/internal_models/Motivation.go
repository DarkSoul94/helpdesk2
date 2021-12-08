package internal_models

import "github.com/shopspring/decimal"

type Motivation struct {
	Support           *MotivSupport
	ByCategory        []*MotivCategory
	TotalTicketsCount uint64
	TotalMotivation   decimal.Decimal
	TotalByShifts     decimal.Decimal
	Total             decimal.Decimal
}

type MotivSupport struct {
	ID    uint64
	Name  string
	Color string
}

type MotivCategory struct {
	ID    uint64
	Name  string
	Count uint64
}

func (c *MotivCategory) countToDecimal() decimal.Decimal {
	return decimal.New(int64(c.Count), 0)
}

func (c *MotivCategory) CalcMotiv(price decimal.Decimal) decimal.Decimal {
	return price.Mul(c.countToDecimal())
}
