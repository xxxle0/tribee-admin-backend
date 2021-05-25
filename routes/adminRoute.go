package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/controllers"
)

func AdminRoutes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1")
	{
		admin := apiV1.Group("/admin")
		{
			admin.GET("/ping", controllers.AdminController)
			admin.POST("/sign-in", controllers.SignIn)
			admin.POST("/sign-up", controllers.SignIn)
		}
	}
}
