package routes

import (
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1")
	{
		admin := apiV1.Group("/admin")
		{
			admin.GET("/ping", AdminAPI.SignIn)
			admin.POST("/sign-in", AdminAPI.SignIn)
			admin.POST("/sign-up", AdminAPI.SignIn)
		}
	}
}
