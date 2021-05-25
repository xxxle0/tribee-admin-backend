package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/services"
)

func AdminController(c *gin.Context) {
	services.AdminSerice()
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

type SignInBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignInController(c *gin.Context) {
	var signInBody SignInBody
	if err := c.ShouldBind(&signInBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.SignIn(signInBody.Email, signInBody.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
