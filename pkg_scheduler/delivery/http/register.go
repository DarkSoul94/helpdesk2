package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_scheduler"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc pkg_scheduler.ISchedulerUsecase, midellwares ...gin.HandlerFunc) {
	h := NewSchedulerHandler(uc)

	tableEndpoints := router.Group("/table")
	tableEndpoints.Use(midellwares...)
	{
		//http://localhost:8888/helpdesk/table/offices_list
		tableEndpoints.POST("/offices_list", h.UpdateOfficesList)
		//http://localhost:8888/helpdesk/table/offices_list
		tableEndpoints.GET("/offices_list", h.GetOfficesList)
		//http://localhost:8888/helpdesk/table/update_schedule
		tableEndpoints.POST("/update_schedule", h.UpdateShiftsShedule)
		//http://localhost:8888/helpdesk/table/schedule?date=2021-10-01
		tableEndpoints.GET("/schedule", h.GetShiftsShedule)

		//http://localhost:8888/helpdesk/table/lateness/check_new
		tableEndpoints.GET("/lateness/check_new", h.CheckNewLateness)
		//http://localhost:8888/helpdesk/table/lateness?date=2021-10-01
		tableEndpoints.GET("/lateness", h.GetSupportLateness)
		//http://localhost:8888/helpdesk/table/lateness/update
		tableEndpoints.POST("/lateness/update", h.UpdateLateness)

		/*TODO вынести в пакет по работе с константами
		//http://localhost:8888/helpdesk/table/lateness_conf
		tableEndpoints.GET("/lateness_conf", h.GetLatenessConf)
		//http://localhost:8888/helpdesk/table/lateness_conf
		tableEndpoints.POST("/lateness_conf", h.SetLatenessConf)*/
	}
}
