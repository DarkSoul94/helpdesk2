package dto

import (
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
	"github.com/shopspring/decimal"
)

type SupportCard struct {
	ID             uint64    `json:"id"`
	Support        *outShort `json:"support"`
	InternalNumber string    `json:"internal_number"`
	MobileNumber   string    `json:"mobile_number"`
	BirthDate      string    `json:"birth_date"`
	IsSenior       bool      `json:"is_senior"`
	Wager          float64   `json:"wager"`
	Senior         *outShort `json:"senior"`
	Comment        string    `json:"comment"`
	Color          string    `json:"color"`
}

type outShort struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func ToOutSupportCard(card *internal_models.Card) SupportCard {
	wager, _ := card.Wager.Float64()
	var outCard SupportCard = SupportCard{
		ID: card.ID,
		Support: &outShort{
			ID:   card.Support.ID,
			Name: card.Support.Name,
		},
		InternalNumber: card.InternalNumber,
		MobileNumber:   card.MobileNumber,
		BirthDate:      card.BirthDate,
		IsSenior:       card.IsSenior,
		Wager:          wager,
		Comment:        card.Comment,
		Color:          card.Color,
	}
	if card.Senior != nil {
		outCard.Senior = &outShort{
			ID:   card.Senior.ID,
			Name: card.Senior.Name,
		}
	}
	return outCard
}

func ToModelSupportCard(inpCard *SupportCard) internal_models.Card {
	var card internal_models.Card = internal_models.Card{
		ID: inpCard.ID,
		Support: &internal_models.Support{
			ID:   inpCard.Support.ID,
			Name: inpCard.Support.Name,
		},
		InternalNumber: inpCard.InternalNumber,
		MobileNumber:   inpCard.MobileNumber,
		BirthDate:      inpCard.BirthDate,
		IsSenior:       inpCard.IsSenior,
		Wager:          decimal.NewFromFloat(inpCard.Wager),
		Comment:        inpCard.Comment,
		Color:          inpCard.Color,
	}
	if inpCard.Senior != nil {
		card.Senior = &internal_models.Support{
			ID:   inpCard.Senior.ID,
			Name: inpCard.Senior.Name,
		}
	}
	return card
}
