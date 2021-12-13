package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_reports"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc pkg_reports.IReportsUsecase, middlewares ...gin.HandlerFunc) {
	h := NewReportsHandler(uc)

	reportsEndpoints := router.Group("/reports")
	reportsEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/reports/motivation?start_date=2021-03-01&end_date=2021-04-01
		reportsEndpoints.GET("/motivation", h.GetMotivation)
		//http://localhost:8888/helpdesk/reports/tickets_status_difference?start_date=2021-03-01&end_date=2021-04-01
		reportsEndpoints.GET("/tickets_status_difference", h.GetTicketStatusDifference)
		//http://localhost:8888/helpdesk/reports/average_grades?start_date=2021-03-01&end_date=2021-04-01
		reportsEndpoints.GET("/average_grades", h.GetAverageGrades)
		//http://localhost:8888/helpdesk/reports/tickets_grades
		reportsEndpoints.POST("/tickets_grades", h.GetTicketsGrades)
		//http://localhost:8888/helpdesk/reports/returned_tickets?start_date=2021-03-01&end_date=2021-04-01
		reportsEndpoints.GET("/returned_tickets", h.GetReturnedTickets)
		//http://localhost:8888/helpdesk/reports/tickets_count?start_date=2021-03-01&end_date=2021-04-01
		reportsEndpoints.GET("/tickets_count", h.GetTicketsCountByDaysHours)
		//http://localhost:8888/helpdesk/reports/supports_statuses?start_date=2021-03-01&end_date=2021-04-01
		reportsEndpoints.GET("/supports_statuses", h.GetSupportsStatusesByWeekDay)
		//http://localhost:8888/helpdesk/reports/supports_shifts?start_date=2021-03-01&end_date=2021-04-01
		reportsEndpoints.GET("/supports_shifts", h.GetSupportsShifts)
		//http://localhost:8888/helpdesk/reports/supports_statuses_history?date=2021-03-01
		reportsEndpoints.GET("/supports_statuses_history", h.GetSupportStatusHistory)
	}
}
