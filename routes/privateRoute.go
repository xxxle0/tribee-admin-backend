package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/controllers"
)

func PrivateRoutes(router *gin.Engine) {
	api := router.Group("/api/private")
	{
		v1 := api.Group("/v1")
		{
			admin := v1.Group("/admin")
			{
				admin.GET("/ping", controllers.AdminController)
				admin.POST("/sign-in", controllers.SignInController)
			}
		}
	}
}
