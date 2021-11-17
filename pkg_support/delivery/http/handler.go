package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc pkg_support.ISupportUsecase
}

// NewHandler ...
func NewHandler(uc pkg_support.ISupportUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) GetSupportList(c *gin.Context) {
	supports, err := h.uc.GetSupportList()
	if err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}

	outSupports := make([]dto.OutUserShort, 0)

	for _, support := range supports {
		outSupports = append(outSupports, dto.ToOutShortSupport(support))
	}
	c.JSON(http.StatusOK, outSupports)
}
