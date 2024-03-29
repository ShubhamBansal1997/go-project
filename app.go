package application

import (
	"go-assignment/config"
	"go-assignment/server"
	"go-assignment/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)
	gin := gin.Default()
	routes.ConfigureRoutes(app, gin)

	err := gin.Run(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
