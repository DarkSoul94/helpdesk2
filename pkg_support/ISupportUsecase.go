package pkg_support

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type ISupportUsecase interface {
	CreateSupport(usersID ...uint64) models.Err
	DeleteSupport(usersID ...uint64) models.Err
	GetSupportList() ([]*internal_models.Support, models.Err)
	GetActiveSupports() ([]*internal_models.Support, models.Err)
	GetStatusesList() ([]*internal_models.Status, models.Err)
	SetSupportStatus(supportID, statusID uint64) models.Err
}
