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
		//http://localhost:5555/helpdesk/support/open_shift
		//http://localhost:5555/helpdesk/support/close_shift
		//http://localhost:5555/helpdesk/support/shift_status
		//http://localhost:5555/helpdesk/support/create_lateness
		//http://localhost:5555/helpdesk/support/active_support_list
		//http://localhost:5555/helpdesk/support/get_current_statuses
		//http://localhost:5555/helpdesk/support/get_support_status
		//http://localhost:5555/helpdesk/support/change_status
	}

	cardEndpoints := router.Group("/support/card")
	cardEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/support/card?id=13
		//http://localhost:5555/helpdesk/support/card/update
		//http://localhost:5555/helpdesk/support/card/seniors
		//http://localhost:5555/helpdesk/support/card/list
	}
}
