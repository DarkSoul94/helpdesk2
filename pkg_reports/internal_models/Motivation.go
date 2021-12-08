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

func Total(suppMotivation []Motivation) Motivation {
	categories := make(map[uint64]*MotivCategory)

	totalMotiv := Motivation{
		Support: &MotivSupport{
			ID:   0,
			Name: "Итого",
		},
		ByCategory:        make([]*MotivCategory, 0),
		TotalTicketsCount: 0,
		TotalMotivation:   decimal.Zero,
		TotalByShifts:     decimal.Zero,
		Total:             decimal.Zero,
	}

	for _, motivation := range suppMotivation {
		totalMotiv.TotalTicketsCount += motivation.TotalTicketsCount
		totalMotiv.TotalMotivation = totalMotiv.TotalMotivation.Add(motivation.TotalMotivation)
		totalMotiv.TotalByShifts = totalMotiv.TotalByShifts.Add(motivation.TotalByShifts)
		totalMotiv.Total = totalMotiv.Total.Add(motivation.Total)
		for _, category := range motivation.ByCategory {
			if cat, ok := categories[category.ID]; ok {
				cat.Count = categories[category.ID].Count + category.Count
			} else {
				categories[category.ID] = category
			}
		}
	}

	for _, category := range categories {
		totalMotiv.ByCategory = append(totalMotiv.ByCategory, category)
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
