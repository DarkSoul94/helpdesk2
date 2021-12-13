package http

import (
	"github.com/DarkSoul94/helpdesk2/pkg_consts"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc pkg_consts.IConstsUsecase, middlewares ...gin.HandlerFunc) {
	h := NewConstsHandler(uc)

	constEndpoints := router.Group("/const")
	constEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/const/:const_name
		constEndpoints.POST("/:const", h.SetConst)
		//http://localhost:8888/helpdesk/const/:const_name
		constEndpoints.GET("/:const", h.GetConst)
	}

	settingsEndpoints := router.Group("/table")
	settingsEndpoints.Use(middlewares...)
	{
		//http://localhost:8888/helpdesk/table/lateness_conf
		settingsEndpoints.POST("/lateness_conf", h.SetSettings)
		//http://localhost:8888/helpdesk/table/lateness_conf
		settingsEndpoints.GET("/lateness_conf", h.GetSettings)
	}
}
