package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/services"
)

type SignInBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminController struct {
	AdminService services.AdminService
}

type AdminControllerI interface {
	SignIn(c *gin.Context)
	Ping(c *gin.Context)
}

func AdminControllerInit(adminService services.AdminService) AdminController {
	return AdminController{AdminService: adminService}
}

func (co *AdminController) SignIn(c *gin.Context) {
	var signInBody SignInBody
	if err := c.ShouldBind(&signInBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := co.AdminService.SignIn(signInBody.Email, signInBody.Password)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (co *AdminController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
