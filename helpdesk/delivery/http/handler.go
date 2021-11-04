package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/helpdesk"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	uc helpdesk.Usecase
}

// NewHandler ...
func NewHandler(uc helpdesk.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// HelloWorld ...
func (h *Handler) HelloWorld(c *gin.Context) {
	h.uc.HelloWorld(c.Request.Context())
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
