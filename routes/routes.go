package routes

import (
	"github.com/pratomoadhi/golden-trail/controller"
	"github.com/pratomoadhi/golden-trail/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/panic", func(c *gin.Context) {
		panic("test panic for Sentry")
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

	transaction := r.Group("/transactions")
	transaction.Use(middleware.JWTAuth())
	{
		transaction.POST("/", controller.CreateTransaction)
		transaction.GET("/", controller.ListTransactions)
	}
}
