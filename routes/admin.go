package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/controllers"
)

func SetupAdminRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			admin := v1.Group("/admin")
			{
				admin.GET("/ping", controllers.AdminController)
			}
		}
	}
}
