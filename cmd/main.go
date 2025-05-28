// @title Golden Trail API
// @version 1.0
// @description API for managing finance tracking
// @host localhost:5000
// @BasePath /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	"github.com/pratomoadhi/golden-trail/config"
	_ "github.com/pratomoadhi/golden-trail/docs"
	"github.com/pratomoadhi/golden-trail/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load config and store it globally
	config.LoadConfig()

	// Use the global config to connect database
	config.ConnectDatabase(config.AppConfig)

	// Start server
	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + config.AppConfig.Port)
}
