package utils

import (
	"github.com/gin-gonic/gin"
)

// APIResponse digunakan untuk membentuk format respons JSON yang konsisten
func APIResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse digunakan untuk mengembalikan pesan error
func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status": status,
		"error":  message,
	})
}
