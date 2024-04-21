package middlewares

import (
	"bytes"
	"io"
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
		// 打印body
		body, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		log.Printf("[LogID:%s] Request: %s", logID, string(body))
		c.Next()

		elapsedTime := time.Since(startTime)

		//如果有错误就打印错误
		log.Printf("[LogID:%s] %s %s %d %v %s",
			logID,
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			elapsedTime,
			c.Request.UserAgent(),
		)
	}
}
