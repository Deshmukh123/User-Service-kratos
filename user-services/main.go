package main

import (
	"user-service/config"
	"user-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":" + cfg.Port)
}
