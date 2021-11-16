package pkg_support

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type ISupportRepo interface {
	CreateSupport(support *internal_models.Support) models.Err
	DeleteSupport(supportID uint64) models.Err
	GetSupport(userID uint64) (*internal_models.Support, models.Err)

	CreateCard(card *internal_models.Card) models.Err
	DeleteCard(supportID uint64) models.Err
	GetCardBySupportID(supportID uint64) (*internal_models.Card, models.Err)
	ResetSenior(seniorID uint64) models.Err

	Close() error
}
