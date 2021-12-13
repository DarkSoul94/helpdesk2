package internal_models

import (
	"sync"

	"github.com/shopspring/decimal"
)

type MotivationByPeriod struct {
	Period      string
	Motivations []*Motivation
}
type Motivation struct {
	Support           *MotivSupport
	ByCategory        []*MotivCategory
	TotalTicketsCount uint64
	TotalMotivation   decimal.Decimal
	TotalByShifts     decimal.Decimal
	Total             decimal.Decimal
}

func NewMotivation(supportID uint64, supportName, color string, motivation decimal.Decimal) *Motivation {
	return &Motivation{
		Support: &MotivSupport{
			ID:    supportID,
			Name:  supportName,
			Color: color,
		},
		ByCategory:        make([]*MotivCategory, 0),
		TotalTicketsCount: 0,
		TotalMotivation:   decimal.Zero,
		TotalByShifts:     motivation,
		Total:             decimal.Zero,
	}
}

func MotivByPeriod(suppMotivation []*Motivation) *Motivation {
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
				categories[category.ID] = &MotivCategory{
					ID:    category.ID,
					Name:  category.Name,
					Count: category.Count,
				}
			}
		}
	}

	for _, category := range categories {
		totalMotiv.ByCategory = append(totalMotiv.ByCategory, category)
	}

	return totalMotiv
}

func CalcMotivByAllPeriod(totalMotiv map[uint64]*Motivation, motivBySupport *Motivation) {
	var (
		mutex      sync.Mutex
		categories = make(map[uint64]uint64)
	)

	mutex.Lock()
	if motiv, ok := totalMotiv[motivBySupport.Support.ID]; ok {
		motiv.TotalTicketsCount += motivBySupport.TotalTicketsCount
		motiv.TotalByShifts = motiv.TotalByShifts.Add(motivBySupport.TotalByShifts)
		motiv.TotalMotivation = motiv.TotalMotivation.Add(motivBySupport.TotalMotivation)
		motiv.Total = motiv.Total.Add(motivBySupport.Total)

		for _, catMotiv := range motivBySupport.ByCategory {
			categories[catMotiv.ID] = catMotiv.Count
		}

		for _, catMotiv := range motiv.ByCategory {
			catMotiv.Count += categories[catMotiv.ID]
		}

	} else {
		suppMotiv := &Motivation{
			Support: &MotivSupport{
				ID:    motivBySupport.Support.ID,
				Name:  motivBySupport.Support.Name,
				Color: motivBySupport.Support.Color,
			},
			TotalTicketsCount: motivBySupport.TotalTicketsCount,
			TotalMotivation:   motivBySupport.TotalMotivation,
			TotalByShifts:     motivBySupport.TotalByShifts,
			Total:             motivBySupport.Total,
		}

		for _, catMotiv := range motivBySupport.ByCategory {
			suppMotiv.ByCategory = append(suppMotiv.ByCategory, &MotivCategory{
				ID:    catMotiv.ID,
				Name:  catMotiv.Name,
				Count: catMotiv.Count,
			})
		}

		totalMotiv[motivBySupport.Support.ID] = suppMotiv
	}
	mutex.Unlock()
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
