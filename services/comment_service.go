package services

import (
	"errors"

	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
	"github.com/google/uuid"
)

// CreateComment adds a new comment to a task
func CreateComment(comment models.Comment) (models.Comment, error) {
	// Ensure TaskID and UserID are valid
	if comment.TaskID == uuid.Nil || comment.UserID == uuid.Nil {
		return models.Comment{}, errors.New("task ID and user ID are required")
	}

	// Save to database
	if err := config.DB.Create(&comment).Error; err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

// GetCommentsByTaskID retrieves all comments for a specific task
func GetCommentsByTaskID(taskID uuid.UUID) ([]models.Comment, error) {
	var comments []models.Comment
	if err := config.DB.Where("task_id = ?", taskID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// DeleteComment removes a comment by its ID
func DeleteComment(commentID uuid.UUID, userID uuid.UUID) error {
	var comment models.Comment

	// Check if comment exists
	if err := config.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return errors.New("comment not found")
	}

	// Ensure only the comment owner can delete it
	if comment.UserID != userID {
		return errors.New("unauthorized to delete this comment")
	}

	// Delete comment
	if err := config.DB.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
