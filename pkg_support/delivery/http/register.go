package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_support"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc pkg_support.ISupportUsecase, middlewares ...gin.HandlerFunc) {
	h := NewHandler(uc)

	supportEndpoints := router.Group("/support")
	supportEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/support/support_list
		supportEndpoints.GET("/support_list", h.GetSupportList)
		//http://localhost:5555/helpdesk/support/status_list
		supportEndpoints.GET("/status_list", h.GetStatusesList)
		//http://localhost:5555/helpdesk/support/open_shift
		supportEndpoints.POST("/open_shift", h.OpenShift)
		//http://localhost:5555/helpdesk/support/create_lateness
		//supportEndpoints.POST("/create_lateness", h.CreateLateness) //TODO Создание опоздания при открытии смены
		//http://localhost:5555/helpdesk/support/close_shift
		supportEndpoints.POST("/close_shift", h.CloseShift)
		//http://localhost:5555/helpdesk/support/shift_status
		supportEndpoints.GET("/shift_status", h.GetShiftStatus)
		//http://localhost:5555/helpdesk/support/active_support_list
		supportEndpoints.GET("/active_support_list", h.GetActiveSupports)
		//http://localhost:5555/helpdesk/support/get_current_statuses
		supportEndpoints.GET("/get_current_statuses", h.GetCurrentStatuses)
		//http://localhost:5555/helpdesk/support/get_support_status
		supportEndpoints.GET("/get_support_status", h.GetSupportStatus)
		//http://localhost:5555/helpdesk/support/change_status
		supportEndpoints.POST("/change_status", h.ChangeSupportStatus)
	}

	cardEndpoints := router.Group("/support/card")
	cardEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/support/card?id=13
		cardEndpoints.GET("", h.GetCard)
		//http://localhost:5555/helpdesk/support/card/update
		cardEndpoints.POST("/update", h.UpdateCard)
		//http://localhost:5555/helpdesk/support/card/seniors
		//http://localhost:5555/helpdesk/support/card/list
	}
}
