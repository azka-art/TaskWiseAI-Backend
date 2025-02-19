package routes

import (
	"github.com/azka-art/taskwise-backend/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes mendaftarkan endpoint autentikasi
func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser) // Register user
		auth.POST("/login", controllers.LoginUser)       // Login user
	}
}
