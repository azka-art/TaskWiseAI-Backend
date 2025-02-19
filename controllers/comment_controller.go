package controllers

import (
	"net/http"

	"github.com/azka-art/taskwise-backend/models"
	"github.com/azka-art/taskwise-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddComment adds a comment to a task
func AddComment(c *gin.Context) {
	// ✅ Convert TaskID from string to uuid.UUID
	taskID, err := uuid.Parse(c.Param("task_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Extract UserID from JWT claims
	userID, exists := c.Get("user_id") // Make sure this is set in JWT middleware
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// ✅ Convert UserID to uuid.UUID
	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	// ✅ Assign TaskID and UserID
	comment.TaskID = taskID
	comment.UserID = userUUID

	newComment, err := services.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, newComment)
}

// GetComments retrieves all comments for a task
func GetComments(c *gin.Context) {
	// ✅ Convert TaskID from string to uuid.UUID
	taskID, err := uuid.Parse(c.Param("task_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	comments, err := services.GetCommentsByTaskID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
