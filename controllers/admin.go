package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/services"
)

func AdminController(c *gin.Context) {
	services.AdminSerice()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SignInController(c *gin.Context) {
	services.SignIn()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
