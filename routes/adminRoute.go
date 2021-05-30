package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/xxxle0/tribee-admin-backend/containers"
	"github.com/xxxle0/tribee-admin-backend/middlewares"
)

func AdminRoutes(router *gin.Engine, Db *sqlx.DB) {
	AdminController := containers.InitializeAdminController()
	apiV1 := router.Group("/api/v1")
	{
		admin := apiV1.Group("/admin")
		{
			admin.GET("/ping", AdminController.Ping)
			admin.POST("/sign-in", AdminController.SignIn)
			admin.GET("/private/ping", middlewares.Authentication, AdminController.Ping)
		}
	}
}
