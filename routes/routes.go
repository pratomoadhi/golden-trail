package routes

import (
	"github.com/pratomoadhi/golden-trail/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		// auth.POST("/login", controller.Login) // Next step
	}
}
