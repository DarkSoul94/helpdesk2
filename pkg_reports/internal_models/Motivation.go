package internal_models

import (
	"github.com/shopspring/decimal"
)

type Motivation struct {
	Support           *MotivSupport
	ByCategory        []MotivCategory
	TotalTicketsCount uint64
	TotalMotivation   decimal.Decimal
	TotalByShifts     decimal.Decimal
	Total             decimal.Decimal
}

func Total(suppMotivation []Motivation) Motivation {
	categories := make(map[uint64]*MotivCategory)
	totalMotiv := NewMotivation(0, "Итого", "", decimal.Zero)

	for _, motivation := range suppMotivation {
		totalMotiv.TotalTicketsCount += motivation.TotalTicketsCount
		totalMotiv.TotalMotivation = totalMotiv.TotalMotivation.Add(motivation.TotalMotivation)
		totalMotiv.TotalByShifts = totalMotiv.TotalByShifts.Add(motivation.TotalByShifts)
		totalMotiv.Total = totalMotiv.Total.Add(motivation.Total)

		for _, category := range motivation.ByCategory {
			if cat, ok := categories[category.ID]; ok {
				cat.Count += category.Count
			} else {
				temp := category
				categories[category.ID] = &temp
			}
		}
	}

	for _, category := range categories {
		totalMotiv.ByCategory = append(totalMotiv.ByCategory, *category)
	}

	return totalMotiv
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

func NewMotivation(supportID uint64, supportName, color string, motivation decimal.Decimal) Motivation {
	return Motivation{
		Support: &MotivSupport{
			ID:    supportID,
			Name:  supportName,
			Color: color,
		},
		ByCategory:        make([]MotivCategory, 0),
		TotalTicketsCount: 0,
		TotalMotivation:   decimal.Zero,
		TotalByShifts:     motivation,
		Total:             decimal.Zero,
	}
}

func SummaryMotiv(c, c2 Motivation) Motivation {
	categories := make(map[uint64]MotivCategory)

	for _, category := range c2.ByCategory {
		if cat, ok := categories[category.ID]; ok {
			cat.Count += category.Count
		} else {
			categories[category.ID] = category
		}
	}

	for index, val := range c.ByCategory {
		c.ByCategory[index].Count = val.Count + categories[val.ID].Count
	}

	c.TotalTicketsCount += c2.TotalTicketsCount
	c.TotalByShifts = c.TotalByShifts.Add(c2.TotalByShifts)
	c.TotalMotivation = c.TotalMotivation.Add(c2.TotalMotivation)
	c.Total = c.Total.Add(c2.Total)

	return c
}
