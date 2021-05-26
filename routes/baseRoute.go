package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/controllers"
)

func BaseRoutes(router *gin.Engine) {
	BaseController := controllers.BaseController{}
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/ping", BaseController.Ping)
	}
}
