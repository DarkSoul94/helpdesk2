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
	startDate := c.Request.URL.Query().Get("start_date")
	endDate := c.Request.URL.Query().Get("end_date")

	motivation, err := h.uc.GetMotivation(startDate, endDate)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outMotivation := make(map[string][]dto.OutMotivation)
	for period, motiv := range motivation {
		outMotivation[period] = dto.ToOutMotivation(motiv)
	}

	c.JSON(http.StatusOK, outMotivation)
}

func (h *ReportsHandler) GetTicketStatusDifference(c *gin.Context) {
	startDate := c.Request.URL.Query().Get("start_date")
	endDate := c.Request.URL.Query().Get("end_date")

	difference, err := h.uc.GetTicketStatusDifference(startDate, endDate)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outDifference := make([]dto.OutTicketStatusDifference, 0)
	for ticket, statuses := range difference {
		outDifference = append(outDifference, dto.ToOutTicketStatusDifference(ticket, statuses))
	}

	c.JSON(http.StatusOK, outDifference)
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
	var param dto.InpParam

	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	grades, err := h.uc.GetTicketsGrade(param.StartDate, param.EndDate, param.UsersID, param.Departments)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outGrades := make([]dto.OutDepartmentTicketGrade, 0)

	for department, departmentGrades := range grades {
		outDepartmentGrades := dto.OutDepartmentTicketGrade{
			Department:             department,
			UsersGrades:            make([]dto.OutUserTicketGrades, 0),
			AvaregeDepartmentGrade: 0,
		}

		for user, userTicketsGrades := range departmentGrades {
			outUserGrades := dto.OutUserTicketGrades{
				UserName:         user,
				TicketsGrades:    make([]dto.OutTicketGrade, 0),
				AverageUserGrade: 0,
			}

			for _, ticketGrade := range userTicketsGrades {
				outGrade := dto.OutTicketGrade{
					TicketID:    ticketGrade.TicketID,
					TicketGrade: ticketGrade.TicketGrade,
				}

				outUserGrades.AverageUserGrade += float64(ticketGrade.TicketGrade)
				outUserGrades.TicketsGrades = append(outUserGrades.TicketsGrades, outGrade)
			}

			outUserGrades.AverageUserGrade = math.Round(outUserGrades.AverageUserGrade/float64(len(outUserGrades.TicketsGrades))*100) / 100
			outDepartmentGrades.AvaregeDepartmentGrade += outUserGrades.AverageUserGrade
			outDepartmentGrades.UsersGrades = append(outDepartmentGrades.UsersGrades, outUserGrades)
		}
		outDepartmentGrades.AvaregeDepartmentGrade = math.Round(outDepartmentGrades.AvaregeDepartmentGrade/float64(len(outDepartmentGrades.UsersGrades))*100) / 100
		outGrades = append(outGrades, outDepartmentGrades)
	}

	c.JSON(http.StatusOK, outGrades)
}

func (h *ReportsHandler) GetReturnedTickets(c *gin.Context) {
	startDate := c.Request.URL.Query().Get("start_date")
	endDate := c.Request.URL.Query().Get("end_date")

	tickets, err := h.uc.GetReturnedTickets(startDate, endDate)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outTickets := make([]dto.OutReturnedTicket, 0)
	for _, ticket := range tickets {
		outTickets = append(outTickets, dto.ToOutReturnedTickets(ticket))
	}

	c.JSON(http.StatusOK, outTickets)
}

func (h *ReportsHandler) GetTicketsCountByDaysHours(c *gin.Context) {
	startDate := c.Request.URL.Query().Get("start_date")
	endDate := c.Request.URL.Query().Get("end_date")

	countByDayHour, err := h.uc.GetTicketsCountByDaysHours(startDate, endDate)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToOutTicketsCountByDay(countByDayHour))
}

func (h *ReportsHandler) GetSupportsStatusesByWeekDay(c *gin.Context) {
	startDate := c.Request.URL.Query().Get("start_date")
	endDate := c.Request.URL.Query().Get("end_date")

	history, err := h.uc.GetSupportsStatusesByWeekDay(startDate, endDate)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToOutSupportStatusesHistory(history))
}

func (h *ReportsHandler) GetSupportsShifts(c *gin.Context) {
	startDate := c.Request.URL.Query().Get("start_date")
	endDate := c.Request.URL.Query().Get("end_date")

	mShifts, err := h.uc.GetSupportsShifts(startDate, endDate)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outShifts := make(map[string][]dto.OutSupportsShift)
	for period, shifts := range mShifts {
		outShifts[period] = make([]dto.OutSupportsShift, 0)
		for _, shift := range shifts {
			outShifts[period] = append(outShifts[period], dto.ToOutSupportShift(shift))
		}
	}

	c.JSON(http.StatusOK, outShifts)
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
