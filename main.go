package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/xxxle0/tribee-admin-backend/routes"
	"github.com/xxxle0/tribee-admin-backend/utils"
)

var Db *sqlx.DB

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("error load config", err)
	}
	Db, err = sqlx.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	routes.SetupRoutes(router, Db)
	router.Run(config.ServerPort)
}
