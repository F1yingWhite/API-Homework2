package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logID := uuid.New().String()
		c.Set("logID", logID)
		startTime := time.Now()

		c.Next()

		elapsedTime := time.Since(startTime)

		log.Printf("[LogID:%s] Method: %s, URI: %s, IP: %s, Status: %d, ElapsedTime: %v, UserAgent: %s, Errors: %v",
			logID,
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			c.Writer.Status(),
			elapsedTime,
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate),
		)
	}
}
