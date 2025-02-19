package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware merekam setiap request yang masuk ke server
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Proses request
		c.Next()

		// Hitung durasi request
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()

		log.Printf("[%d] %s %s | %v", statusCode, c.Request.Method, c.Request.URL.Path, latency)
	}
}
