package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc pkg_ticket.ITicketUsecase, middlewares ...gin.HandlerFunc) {
	h := NewTicketHandler(uc)

	categoryEndpoints := router.Group("/category")
	categoryEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/category/create
		categoryEndpoints.POST("/create", h.CreateCategory)
	}
}
