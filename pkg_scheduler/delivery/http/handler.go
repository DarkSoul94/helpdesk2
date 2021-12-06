package http

import (
	"net/http"
	"time"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler/internal_models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc pkg_scheduler.ISchedulerUsecase
}

// NewHandler ...
func NewSchedulerHandler(uc pkg_scheduler.ISchedulerUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) UpdateOfficesList(c *gin.Context) {
	var offices []dto.OutOffice
	if err := c.BindJSON(&offices); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	if err := h.uc.UpdateOfficesList(dto.ToModelOffices(offices)); err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) GetOfficesList(c *gin.Context) {
	mActual, mDeleted, err := h.uc.GetOfficesList()
	if err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "actual": dto.ToOutOffices(mActual), "deleted": dto.ToOutOffices(mDeleted)})
}

func (h *Handler) UpdateShiftsShedule(c *gin.Context) {
	var (
		schedule  []dto.OutScheduleCell
		mSchedule []*internal_models.Cell
	)

	if err := c.BindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	for _, cell := range schedule {
		mSchedule = append(mSchedule, dto.ToModelScheduleCell(cell))
	}
	if err := h.uc.UpdateShiftsSchedule(mSchedule); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) GetShiftsShedule(c *gin.Context) {
	//TODO добавить получение групп из пакета саппортов
}

func (h *Handler) CheckNewLateness(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "exist": h.uc.CheckNewLateness()})
}

func (h *Handler) GetSupportLateness(c *gin.Context) {
	var (
		date    string
		list    []*internal_models.Lateness
		outList = make([]*dto.OutSupportLateness, 0)

		err models.Err
	)
	if date = c.Request.URL.Query().Get("date"); len(date) == 0 {
		date = time.Now().Local().String()
	}

	if list, err = h.uc.GetLateness(date); err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}
	for _, record := range list {
		outList = append(outList, dto.ToOutSupportLateness(record))
	}
	mDecisions := internal_models.GetDicisionsList()
	c.JSON(http.StatusOK,
		map[string]interface{}{
			"status":    "ok",
			"desicions": dto.ToOutDecisionsList(mDecisions),
			"lateness":  outList,
		})
}

func (h *Handler) UpdateLateness(c *gin.Context) {
	type updLateness struct {
		LatenessID uint64 `json:"lateness_id"`
		DecisionID uint64 `json:"decision_id"`
	}
	var updLate updLateness

	if err := c.BindJSON(&updLate); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}
	if err := h.uc.UpdateLateness(updLate.LatenessID, updLate.DecisionID); err != nil {
		c.JSON(err.Code(), map[string]string{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
