package http

import (
	"github.com/DarkSoul94/helpdesk2/auth"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, ucAuth auth.AuthUC) {
	h := NewHandler(ucAuth)

	authEndpoints := router.Group("/auth")
	{
		//http://localhost:5555/helpdesk/auth/signin
		authEndpoints.POST("/signin", h.SignIn)

	}
}
