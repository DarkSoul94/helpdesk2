package internal_models

import (
	"github.com/shopspring/decimal"
)

type Card struct {
	ID             uint64
	Support        *Support
	InternalNumber string
	MobileNumber   string
	BirthDate      string
	IsSenior       bool
	Senior         *Support
	Wager          decimal.Decimal
	Comment        string
	Color          string
}

//NewSupportCard создает объект карточки суппорта без старшего и с установленнм белым цветом
func NewSupportCard(userID uint64) *Card {
	return &Card{
		Support: &Support{
			ID: userID,
		},
		InternalNumber: "",
		MobileNumber:   "",
		BirthDate:      "",
		Wager:          decimal.New(0, 2),
		Comment:        "",
		Color:          "#FFFFFF",
	}
}
