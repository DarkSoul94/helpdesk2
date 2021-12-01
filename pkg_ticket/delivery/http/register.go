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
		//http://localhost:8888/helpdesk/ticket_status/history?ticket_id=23
		ticketStatusEndpoints.GET("/history", h.GetTicketStatusHistory)
	}

	ticketEndpoints := router.Group("/ticket")
	ticketEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/ticket/create
		ticketEndpoints.POST("/create", h.CreateTicket)
		//http://localhost:8888/helpdesk/ticket/generate_tickets
		ticketEndpoints.POST("generate_tickets", h.GenerateTickets)
		//http://localhost:8888/helpdesk/ticket/update
		ticketEndpoints.POST("/update", h.UpdateTicket)
		//http://localhost:8888/helpdesk/ticket/tickets_list?count=5&offset=0
		ticketEndpoints.GET("/tickets_list", h.GetTicketsList)
		//http://localhost:8888/helpdesk/ticket/ticket?ticket_id=23
		ticketEndpoints.GET("/ticket", h.GetTicket)
		//http://localhost:8888/helpdesk/ticket/steal
		ticketEndpoints.POST("/steal", h.StealTicket)
		//http://localhost:8888/helpdesk/ticket/ticket_grade
		ticketEndpoints.POST("/ticket_grade", h.TicketGrade)
	}

	commentEndpoints := router.Group("/comment")
	commentEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/comment/create
		commentEndpoints.POST("/create", h.CreateComment)
	}

	resolveTicketEndpoints := router.Group("/resolve_ticket")
	resolveTicketEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/resolve_ticket/check_exist
		resolveTicketEndpoints.GET("/check_exist", h.CheckNeedApprovalTicketExist)
		//http://localhost:8888/helpdesk/resolve_ticket/resolve_tickets_list?count=5&offset=0
		resolveTicketEndpoints.GET("/resolve_tickets_list", h.GetApprovalTicketList)
		//http://localhost:8888/helpdesk/resolve_ticket/resolve
		resolveTicketEndpoints.POST("/resolve", h.ResolveTicket)
	}

	fileEndpoints := router.Group("/file")
	fileEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/file/?file_id=23
		fileEndpoints.GET("/", h.GetFile)
	}

}
