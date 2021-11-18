package dto

import (
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/shopspring/decimal"
)

type InpCategory struct {
	ID          uint64  `json:"category_id"`
	Name        string  `json:"category_name"`
	Significant bool    `json:"significant"`
	Old         bool    `json:"old_category"`
	Price       float64 `json:"price"`
}

type OutCategoryWithSections struct {
	ID          uint64               `json:"category_id"`
	Name        string               `json:"category_name"`
	Significant bool                 `json:"significant"`
	Old         bool                 `json:"old_category"`
	Price       float64              `json:"price"`
	Sections    []OutCategorySection `json:"sections"`
}

func ToModelCategory(cat InpCategory) *internal_models.Category {
	return &internal_models.Category{
		ID:          cat.ID,
		Name:        cat.Name,
		Significant: cat.Significant,
		Old:         cat.Old,
		Price:       decimal.NewFromFloat(cat.Price),
	}
}

func ToOutCategory(cat *internal_models.Category) InpCategory {
	fPrice, _ := cat.Price.Truncate(2).Float64()
	return InpCategory{
		ID:          cat.ID,
		Name:        cat.Name,
		Significant: cat.Significant,
		Old:         cat.Old,
		Price:       fPrice,
	}
}

func ToOutCategoryWithSections(list internal_models.CategorySectionList) OutCategoryWithSections {
	fPrice, _ := list.Category.Price.Truncate(2).Float64()
	outList := OutCategoryWithSections{
		ID:          list.Category.ID,
		Name:        list.Category.Name,
		Significant: list.Category.Significant,
		Old:         list.Category.Old,
		Price:       fPrice,
	}

	for _, outSec := range list.Sections {
		outList.Sections = append(outList.Sections, ToOutCategorySection(outSec))
	}

	return outList
}
