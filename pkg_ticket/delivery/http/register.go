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
		//http://localhost:8888/helpdesk/category/update
		categoryEndpoints.POST("/update", h.UpdateCategory)
	}

	sectionEndpoints := router.Group("/section")
	sectionEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/section/create
		sectionEndpoints.POST("/create", h.CreateCategorySection)
		//http://localhost:8888/helpdesk/section/update
		sectionEndpoints.POST("/update", h.UpdateCategorySection)
		//http://localhost:8888/helpdesk/section/
		sectionEndpoints.GET("/", h.GetCategorySection)
		//http://localhost:8888/helpdesk/section/section_list
		sectionEndpoints.GET("/section_list", h.GetCategorySectionList)
	}

}
