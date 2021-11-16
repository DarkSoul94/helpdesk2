package internal_models

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/shopspring/decimal"
)

type Card struct {
	ID             uint64
	Support        *models.User
	InternalNumber string
	MobileNumber   string
	BirthDate      string
	IsSenior       bool
	Senior         *models.User
	Wager          decimal.Decimal
	Comment        string
	Color          string
}

//NewSupportCard создает объект карточки суппорта без старшего и с установленнм белым цветом
func NewSupportCard(userID uint64) *Card {
	return &Card{
		Support: &models.User{
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
