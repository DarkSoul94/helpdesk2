package dto

import "github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"

type SupportCard struct {
	ID             uint64
	Support        *OutUserShort
	InternalNumber string
	MobileNumber   string
	BirthDate      string
	IsSenior       bool
	Wager          float64
	Senior         *OutUserShort
	Comment        string
	Color          string
}

func ToOutSupportCard(card *internal_models.Card) SupportCard {
	wager, _ := card.Wager.Float64()
	var outCard SupportCard = SupportCard{
		ID: card.ID,
		Support: &OutUserShort{
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
		outCard.Senior = &OutUserShort{
			ID:   card.Senior.ID,
			Name: card.Senior.Name,
		}
	}
	return outCard
}
