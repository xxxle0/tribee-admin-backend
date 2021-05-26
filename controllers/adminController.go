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

func AdminControllerInit(adminService services.AdminService) AdminController {
	return AdminController{AdminService: adminService}
}

func (adminController *AdminController) SignIn(c *gin.Context) {
	var signInBody SignInBody
	if err := c.ShouldBind(&signInBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adminController.AdminService.SignIn(signInBody.Email, signInBody.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}