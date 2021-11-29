package dto

import (
	"regexp"
	"strings"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
	"github.com/shopspring/decimal"
)

type SupportCard struct {
	ID             uint64    `json:"id"`
	Support        *OutShort `json:"support"`
	InternalNumber string    `json:"internal_number"`
	MobileNumber   string    `json:"mobile_number"`
	BirthDate      string    `json:"birth_date"`
	IsSenior       bool      `json:"is_senior"`
	Wager          float64   `json:"wager"`
	Senior         *OutShort `json:"senior"`
	Comment        string    `json:"comment"`
	Color          string    `json:"color"`
}

type SupportCardShort struct {
	ID       uint64    `json:"id"`
	Support  *OutShort `json:"support"`
	IsSenior bool      `json:"is_senior"`
	Senior   *OutShort `json:"senior"`
	Color    string    `json:"color"`
}

func ToOutSupportCardShort(card *internal_models.Card) SupportCardShort {
	outCard := SupportCardShort{
		ID: card.ID,
		Support: &OutShort{
			ID:   card.Support.ID,
			Name: card.Support.Name,
		},
		IsSenior: card.IsSenior,
		Color:    card.Color,
	}
	if card.Senior != nil {
		outCard.Senior = &OutShort{
			ID:   card.Senior.ID,
			Name: card.Senior.Name,
		}
	}
	return outCard
}

func ToOutSupportCard(card *internal_models.Card) SupportCard {
	wager, _ := card.Wager.Float64()
	var outCard SupportCard = SupportCard{
		ID: card.ID,
		Support: &OutShort{
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
		outCard.Senior = &OutShort{
			ID:   card.Senior.ID,
			Name: card.Senior.Name,
		}
	}
	return outCard
}

func ToModelSupportCard(inpCard *SupportCard) *internal_models.Card {
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
	return &card
}

func (sc *SupportCard) ValidateCard() models.Err {
	var (
		errArray []string = make([]string, 0)
		err      error
	)
	if err = sc.validateBirthDate(); err != nil {
		errArray = append(errArray, err.Error())
	}
	if err = sc.validateMobileNumber(); err != nil {
		errArray = append(errArray, err.Error())
	}
	if len(errArray) > 0 {
		result := strings.Join(errArray, "; ")
		return models.BadRequest(result)
	}
	return nil
}

func (sc *SupportCard) validateBirthDate() models.Err {
	if len(sc.BirthDate) > 0 {
		query := `(?:[1-9]|[12][0-9]|3[01])\.(?:[1-9]|1[012])\.(?:(19|20|21)\d\d)`
		re := regexp.MustCompile(query)
		if !re.MatchString(sc.BirthDate) {
			return models.BadRequest("Дата рождения не удовлетворяет шаблон ДД.ММ.ГГГГ")
		}
	}
	return nil
}

func (sc *SupportCard) validateMobileNumber() models.Err {
	if len(sc.MobileNumber) > 0 {
		query := `(^\+?3?8?(0\d{9})$)|(^0[\d\W]{10,13})`
		re := regexp.MustCompile(query)
		if !re.MatchString(sc.MobileNumber) {
			return models.BadRequest("Неверный формат мобильного номера телефона")
		}
	}
	return nil
}
