package pkg_support

import "github.com/DarkSoul94/helpdesk2/models"

type ISupportUsecase interface {
	CreateSupport(usersID ...uint64) models.Err
	DeleteSupport(usersID ...uint64) models.Err
}
