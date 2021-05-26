package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/xxxle0/tribee-admin-backend/controllers"
	"github.com/xxxle0/tribee-admin-backend/services"
)

func InitAdminAPI() *controllers.AdminController {
	panic(wire.Build(controllers.AdminControllerInit, services.AdminServiceInit))
}

func AdminRoutes(router *gin.Engine) {
	AdminAPI := InitAdminAPI()
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
