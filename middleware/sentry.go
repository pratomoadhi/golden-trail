package middleware

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SentryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				hub := sentry.CurrentHub().Clone()
				hub.Scope().SetRequest(c.Request)
				hub.Recover(err)
				hub.Flush(time.Second * 2)
				c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			}
		}()
		c.Next()
	}
}
