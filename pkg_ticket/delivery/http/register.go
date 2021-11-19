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

	regionEndpoints := router.Group("/region")
	regionEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/region/create
		regionEndpoints.POST("/create", h.CreateRegion)
		//http://localhost:8888/helpdesk/region/?region_id=23
		regionEndpoints.DELETE("/", h.DeleteRegion)
		//http://localhost:8888/helpdesk/region/update
		regionEndpoints.POST("/update", h.UpdateRegion)
	}

	filialEndpoints := router.Group("/filial")
	filialEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/filial/create
		filialEndpoints.POST("/create", h.CreateFilial)
		//http://localhost:8888/helpdesk/filial/update
		filialEndpoints.POST("/update", h.UpdateFilial)
		//http://localhost:8888/helpdesk/filial/?filial_id=23
		filialEndpoints.DELETE("/", h.DeleteFilial)
		//http://localhost:8888/helpdesk/filial/filial_list
		filialEndpoints.GET("/filial_list", h.GetFilialList)
	}

	ticketStatusEndpoints := router.Group("/ticket_status")
	ticketStatusEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/ticket_status/
		ticketStatusEndpoints.GET("/", h.GetTicketStatuses)
		//http://localhost:8888/helpdesk/ticket_status/list
		ticketStatusEndpoints.GET("/list", h.GetAllTicketStatuses)
	}
}