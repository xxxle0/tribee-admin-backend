package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(route *gin.Engine) string {
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
