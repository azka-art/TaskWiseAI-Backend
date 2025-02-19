package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware mengizinkan frontend mengakses API dari domain berbeda
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Ganti dengan domain frontend
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		// Jika OPTIONS request (Pre-flight), langsung di-respond tanpa lanjut ke handler berikutnya
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
