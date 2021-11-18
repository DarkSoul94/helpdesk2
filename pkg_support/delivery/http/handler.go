package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_support"
)

type Handler struct {
	ucUserManager pkg_support.ISupportUsecase
}

// NewHandler ...
func NewHandler(uc pkg_support.ISupportUsecase) *Handler {
	return &Handler{
		ucUserManager: uc,
	}
}
