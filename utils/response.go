package utils

import (
	"github.com/gin-gonic/gin"
)

// Response represents a standardized API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// APIResponse returns a success response with data
func APIResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse returns an error response
func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Success: false,
		Error:   message,
	})
}

// NotFoundResponse returns a standardized not found response
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, 404, message)
}

// ValidationErrorResponse returns a standardized validation error response
func ValidationErrorResponse(c *gin.Context, message string, errors interface{}) {
	c.JSON(400, Response{
		Success: false,
		Error:   message,
		Data:    errors, // Include validation error details
	})
}
