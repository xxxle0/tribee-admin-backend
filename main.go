package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/xxxle0/tribee-admin-backend/repositories"
	"github.com/xxxle0/tribee-admin-backend/routes"
	"github.com/xxxle0/tribee-admin-backend/utils"
)

type DB struct {
	db *sqlx.DB
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("error load config", err)
	}
	repositories.Db, err = sqlx.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
