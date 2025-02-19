package routes

import (
	"github.com/azka-art/taskwise-backend/controllers"
	"github.com/azka-art/taskwise-backend/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers all API routes
func SetupRoutes(router *gin.Engine) {
	// Public Routes (No Authentication Required)
	router.POST("/api/register", controllers.RegisterUser)
	router.POST("/api/login", controllers.LoginUser)

	// Protected Routes with JWT Middleware
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
	}

	// Register Task Routes (Inside Protected API)
	RegisterTaskRoutes(protected)
}
