package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	uc pkg_ticket.ITicketUsecase
}

func NewTicketHandler(uc pkg_ticket.ITicketUsecase) *TicketHandler {
	return &TicketHandler{
		uc: uc,
	}
}

func (h *TicketHandler) CreateCategory(c *gin.Context) {
	var cat dto.InpCategory

	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.CreateCategory(dto.ToModelCategory(cat))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "category_id": id})
}
