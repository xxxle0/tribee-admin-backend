package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/xxxle0/tribee-admin-backend/controllers"
	"github.com/xxxle0/tribee-admin-backend/middlewares"
	"github.com/xxxle0/tribee-admin-backend/repositories"
	"github.com/xxxle0/tribee-admin-backend/services"
)

func AdminRoutes(router *gin.Engine, Db *sqlx.DB) {
	UserRepository := repositories.UserRepositoryInit(Db)
	AdminService := services.AdminServiceInit(UserRepository)
	AdminController := controllers.AdminControllerInit(AdminService)

	apiV1 := router.Group("/api/v1")
	{
		admin := apiV1.Group("/admin")
		{
			admin.GET("/ping", AdminController.Ping)
			admin.POST("/sign-in", AdminController.SignIn)
			admin.POST("/sign-up", AdminController.SignUp)
			admin.GET("/private/ping", middlewares.Authentication(), AdminController.Ping)
		}
	}
}
