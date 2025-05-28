package main

import (
	"github.com/yourusername/golden-trail/config"
	"github.com/yourusername/golden-trail/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":" + cfg.Port)
}
