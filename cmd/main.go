// @title Golden Trail API
// @version 1.0
// @description API for managing finance tracking
// @host localhost:8080
// @BasePath /
package main

import (
	"github.com/pratomoadhi/golden-trail/config"
	"github.com/pratomoadhi/golden-trail/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/pratomoadhi/golden-trail/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.LoadConfig()
	config.ConnectDatabase(cfg)

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + cfg.Port)
}
