package repositories

import (
	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
)

// CreateComment menambahkan komentar ke database
func CreateComment(comment *models.Comment) error {
	return config.DB.Create(comment).Error
}

// GetCommentsByTaskID mengambil semua komentar berdasarkan Task ID
func GetCommentsByTaskID(taskID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := config.DB.Where("task_id = ?", taskID).Find(&comments).Error
	return comments, err
}
