package http

import (
	"math"
	"net/http"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/pkg_reports"
	"github.com/gin-gonic/gin"
)

type ReportsHandler struct {
	uc pkg_reports.IReportsUsecase
}

func NewReportsHandler(uc pkg_reports.IReportsUsecase) *ReportsHandler {
	return &ReportsHandler{
		uc: uc,
	}
}

func (h *ReportsHandler) GetMotivation(c *gin.Context) {

}

func (h *ReportsHandler) GetTicketStatusDifference(c *gin.Context) {

}

func (h *ReportsHandler) GetAverageGrades(c *gin.Context) {
	var (
		outAVG     []dto.OutAverageGrade
		gradebyAll float64 = 0
	)
	startDate := c.Request.URL.Query().Get("start_date")
	endDate := c.Request.URL.Query().Get("end_date")

	grades, err := h.uc.GetAverageGradesBySupport(startDate, endDate)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	for name, grade := range grades {
		outAVG = append(outAVG, dto.OutAverageGrade{
			Name:         name,
			AverageGrade: grade,
		})
		gradebyAll += grade
	}

	count := len(grades)
	if count != 0 {
		AllAvg := gradebyAll / float64(count)
		outAVG = append(outAVG, dto.OutAverageGrade{
			Name:         "Отдел ТП",
			AverageGrade: math.Round(AllAvg*100) / 100,
		})
	}

	c.JSON(http.StatusOK, outAVG)
}

func (h *ReportsHandler) GetTicketsGrades(c *gin.Context) {

}

func (h *ReportsHandler) GetReturnedTickets(c *gin.Context) {

}

func (h *ReportsHandler) GetTicketsCountByDaysHours(c *gin.Context) {

}

func (h *ReportsHandler) GetSupportsStatusesByWeekDay(c *gin.Context) {

}

func (h *ReportsHandler) GetSupportsShifts(c *gin.Context) {

}

func (h *ReportsHandler) GetSupportStatusHistory(c *gin.Context) {
	date := c.Request.URL.Query().Get("date")

	historyList, err := h.uc.GetSupportsStatusHistory(date)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToOutSupportStatusHistoryList(historyList))
}
