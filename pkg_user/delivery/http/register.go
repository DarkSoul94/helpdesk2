package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc pkg_user.UserManagerUC, middlewares ...gin.HandlerFunc) {
	h := NewHandler(uc)

	usersEndpoints := router.Group("/user")
	usersEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/user/list
		usersEndpoints.GET("/list", h.GetUsersList)
		//http://localhost:5555/helpdesk/user/update
		usersEndpoints.POST("/update", h.UpdateUser)
	}

	groupsEndpoints := router.Group("/group")
	groupsEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/group/create
		groupsEndpoints.POST("/create", h.CreateGroup)
		//http://localhost:5555/helpdesk/group/update
		groupsEndpoints.POST("/update", h.UpdateGroup)
		//http://localhost:5555/helpdesk/group/list
		groupsEndpoints.GET("/list", h.GetGroupsList)
		//http://localhost:5555/helpdesk/group/1
		groupsEndpoints.GET("/:id", h.GetGroup)
	}
	/*
		permissionsEndpoints := router.Group("/permissions")
		permissionsEndpoints.Use(middlewares...)
		{
			//http://localhost:5555/helpdesk/permissions/list
			permissionsEndpoints.GET("/list", h.GetPermList)
		}
	*/
}
