package dto

import (
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/cat_sec_manager"
	"github.com/shopspring/decimal"
)

type InpCategory struct {
	CategoryID          uint64  `json:"category_id"`
	CategoryName        string  `json:"category_name"`
	SignificantCategory bool    `json:"significant"`
	OldCategory         bool    `json:"old_category"`
	Price               float64 `json:"price"`
}

func ToModelCategory(cat InpCategory) *cat_sec_manager.Category {
	return &cat_sec_manager.Category{
		CategoryID:          cat.CategoryID,
		CategoryName:        cat.CategoryName,
		SignificantCategory: cat.SignificantCategory,
		OldCategory:         cat.OldCategory,
		Price:               decimal.NewFromFloat(cat.Price),
	}
}
