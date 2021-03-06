package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc pkg_user.IUserUsecase, middlewares ...gin.HandlerFunc) {
	h := NewHandler(uc)

	usersEndpoints := router.Group("/user")
	usersEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/user/
		usersEndpoints.GET("/", h.GetUsersList)
		//http://localhost:5555/helpdesk/user/update
		usersEndpoints.POST("/update", h.UpdateUser)
		//http://localhost:5555/helpdesk/user/departments_list
		usersEndpoints.GET("/departments_list", h.GetDepartmentsList)
	}

	groupsEndpoints := router.Group("/group")
	groupsEndpoints.Use(middlewares...)
	{
		//http://localhost:5555/helpdesk/group/create
		groupsEndpoints.POST("/create", h.CreateGroup)
		//http://localhost:5555/helpdesk/group/update
		groupsEndpoints.POST("/update", h.UpdateGroup)
		//http://localhost:5555/helpdesk/group/
		groupsEndpoints.GET("/", h.GetGroupsList)
		//http://localhost:5555/helpdesk/group/for_resolve
		groupsEndpoints.GET("/for_resolve", h.GetGroupsListForResolve)
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
