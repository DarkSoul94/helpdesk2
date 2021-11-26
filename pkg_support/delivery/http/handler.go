package http

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetActiveSupports(c *gin.Context) {
	supports, err := h.uc.GetActiveSupports()
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

func (h *Handler) OpenShift(c *gin.Context) {
	type support struct {
		SupportID uint64 `json:"support_id"`
	}
	var sup support
	if err := c.BindJSON(&sup); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	user, _ := c.Get(global_const.CtxUserKey)

	if err := h.uc.OpenShift(sup.SupportID, user.(*models.User)); err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) CloseShift(c *gin.Context) {
	type support struct {
		SupportID uint64 `json:"support_id"`
	}
	var sup support
	if err := c.BindJSON(&sup); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}
	user, _ := c.Get(global_const.CtxUserKey)

	if err := h.uc.CloseShift(sup.SupportID, user.(*models.User)); err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetShiftStatus(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)
	shift, err := h.uc.GetLastShift(user.(*models.User).ID)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "shift_status": false})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "shift_status": !shift.ClosingStatus})
}

func (h *Handler) GetCurrentStatuses(c *gin.Context) {
	statuses, total, err := h.uc.GetCurrentStatuses()
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToOutCurrentStatus(statuses, total))
}

func (h *Handler) GetSupportStatus(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)
	status, err := h.uc.GetSupportStatus(user.(*models.User).ID)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToOutSupportStatus(status))
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

func (h *Handler) GetCard(c *gin.Context) {
	cardID, _ := strconv.ParseUint(c.Request.URL.Query().Get("id"), 10, 64)
	card, fullErr := h.uc.GetCard(cardID)
	if fullErr != nil {
		c.JSON(fullErr.Code(), map[string]string{"status": "error", "error": fullErr.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToOutSupportCard(card))
}

func (h *Handler) UpdateCard(c *gin.Context) {
	var (
		card dto.SupportCard
	)
	if err := c.BindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}
	if err := card.ValidateCard(); err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}
	if err := h.uc.UpdateCard(dto.ToModelSupportCard(&card)); err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})

}
