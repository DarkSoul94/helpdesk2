package pkg_support

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type ISupportRepo interface {
	CreateSupport(support *internal_models.Support) models.Err
	DeleteSupport(supportID uint64) models.Err
	UpdateSupport(support *internal_models.Support) models.Err
	GetSupport(userID uint64) (*internal_models.Support, models.Err)
	GetSupportList() ([]*internal_models.Support, models.Err)
	GetSupportListForToday() ([]*internal_models.Support, models.Err)
	GetActiveSupports() ([]*internal_models.Support, models.Err)
	GetRandomFreeSupport() (*internal_models.Support, models.Err)
	GetPrioritizedSupportID() uint64

	UpdateShift(shift *internal_models.Shift) models.Err

	CheckForActivity(supportID uint64) bool
	CheckForBusy(supportID uint64) bool
	CreateSupportActivity(supportID, ticketID uint64) models.Err
	RemoveSupportActivity(ticketID uint64) models.Err
	UpdateSupportActivity(supportID, ticketID uint64) models.Err

	GetStatus(statusID uint64) (*internal_models.Status, models.Err)
	GetStatusesList() ([]*internal_models.Status, models.Err)

	GetLastShift(supportID uint64) (*internal_models.Shift, models.Err)

	CreateHistoryRecord(statHistory *internal_models.StatusHistory) models.Err
	UpdateHistoryRecord(statHistory *internal_models.StatusHistory) models.Err
	GetLastStatusHistory(supportID, shiftID uint64) (*internal_models.StatusHistory, models.Err)

	CreateCard(card *internal_models.Card) models.Err
	DeleteCard(supportID uint64) models.Err
	GetCardBySupportID(supportID uint64) (*internal_models.Card, models.Err)
	ResetSenior(seniorID uint64) models.Err

	Close() error
}
