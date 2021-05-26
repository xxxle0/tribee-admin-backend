package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (co *BaseController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
