package controllers

import (
	"net/http"

	"github.com/azka-art/taskwise-backend/services"
	"github.com/gin-gonic/gin"
)

// PredictTaskPriority calls AI service to get task priority
func PredictTaskPriority(c *gin.Context) {
	var input struct {
		PriorityLevel     int     `json:"priority_level"`
		DaysUntilDeadline float32 `json:"days_until_deadline"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	predictedPriority, err := services.GetAIPrediction(input.PriorityLevel, input.DaysUntilDeadline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI service error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"predicted_priority": predictedPriority})
}
