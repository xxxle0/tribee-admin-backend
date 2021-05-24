package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xxxle0/tribee-admin-backend/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
