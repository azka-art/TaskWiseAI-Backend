package routes

import (
	"github.com/azka-art/taskwise-backend/controllers"
	"github.com/gin-gonic/gin"
)

// AIRoutes mendaftarkan endpoint untuk AI
func AIRoutes(router *gin.Engine) {
	ai := router.Group("/ai")
	{
		ai.POST("/predict", controllers.PredictTaskPriority) // Prediksi prioritas tugas
	}
}
