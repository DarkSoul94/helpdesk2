package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
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

func (h *Handler) GetStatusesList(c *gin.Context) {
	statuses, err := h.uc.GetStatusesList()
	if err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}
	outStatuses := make([]dto.OutSupportStatus, 0)
	for _, status := range statuses {
		outStatuses = append(outStatuses, dto.ToOutSupportStatus(status))
	}
	c.JSON(http.StatusOK, outStatuses)
}

//ChangeSupportStatus ...
func (h *Handler) ChangeSupportStatus(c *gin.Context) {
	type inpSupport struct {
		SupportID uint64 `json:"support_id"`
		StatusID  uint64 `json:"support_status_id"`
	}
	var support inpSupport
	if err := c.BindJSON(&support); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}
	if support.SupportID == 0 {
		user, _ := c.Get(global_const.CtxUserKey)
		support.SupportID = user.(*models.User).ID
	}

	if err := h.uc.SetSupportStatus(support.SupportID, support.StatusID); err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}