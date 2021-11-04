package http

import (
	"github.com/DarkSoul94/helpdesk2/helpdesk"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, uc helpdesk.Usecase) {
	h := NewHandler(uc)

	apiEndpoints := router.Group("/api")
	{
		apiEndpoints.POST("/", h.HelloWorld)
	}
}
