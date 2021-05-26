package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(router *gin.Engine, Db *sqlx.DB) {
	AdminRoutes(router, Db)
}
