package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{

		v1 := api.Group("/v1")
		{
			admin := v1.Group("/admin")
			{
				admin.GET("/ping", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"message": "pong",
					})
				})
			}

		}
	}
}
