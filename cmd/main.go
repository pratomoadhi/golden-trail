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
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pratomoadhi/golden-trail/config"
	_ "github.com/pratomoadhi/golden-trail/docs"
	"github.com/pratomoadhi/golden-trail/middleware"
	"github.com/pratomoadhi/golden-trail/routes"
)

func main() {
	// Load config and connect database
	config.LoadConfig()
	config.ConnectDatabase(config.AppConfig)

	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.AppConfig.SentryDsn,
		TracesSampleRate: 1.0,
		Environment:      "production",
	})
	if err != nil {
		log.Fatalf("Sentry init failed: %v", err)
	}
	defer sentry.Flush(2 * time.Second)

	r := gin.Default()

	// Attach Sentry middleware
	r.Use(middleware.SentryMiddleware())

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
