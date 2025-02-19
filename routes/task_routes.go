package routes

import (
	"github.com/azka-art/taskwise-backend/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterTaskRoutes sets up task-related routes
func RegisterTaskRoutes(router *gin.RouterGroup) {
	tasks := router.Group("/tasks") // ✅ Uses the protected API group
	{
		tasks.POST("/", controllers.CreateTask)
		tasks.GET("/", controllers.GetTasks)
		tasks.PUT("/:id", controllers.UpdateTask)
		tasks.DELETE("/:id", controllers.DeleteTask)
	}
}
