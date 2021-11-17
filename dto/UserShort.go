package dto

import (
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_support/internal_models"
)

type OutUserShort struct {
	ID   uint64 `json:"user_id"`
	Name string `json:"user_name"`
}

func ToOutShortUser(user *models.User) OutUserShort {
	return OutUserShort{
		ID:   user.ID,
		Name: user.Name,
	}
}

func ToOutShortSupport(support *internal_models.Support) OutUserShort {
	return OutUserShort{
		ID:   support.ID,
		Name: support.Name,
	}
}
