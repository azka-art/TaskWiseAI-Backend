package services

import (
	"errors"

	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
	"github.com/google/uuid"
)

// CreateTask creates a new task
func CreateTask(task models.Task) (models.Task, error) {
	// Ensure task ID is generated if not provided
	if task.ID == uuid.Nil {
		task.ID = uuid.New()
	}

	if err := config.DB.Create(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// GetAllTasks retrieves all tasks
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := config.DB.Preload("Comments").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// ✅ Fix: UpdateTask now accepts `uuid.UUID`
func UpdateTask(id uuid.UUID, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	if err := config.DB.First(&task, "id = ?", id).Error; err != nil {
		return models.Task{}, errors.New("task not found")
	}

	// Update fields only if new values are provided
	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}
	if updatedTask.Priority != "" {
		task.Priority = updatedTask.Priority
	}
	if updatedTask.Status != "" {
		task.Status = updatedTask.Status
	}
	if updatedTask.Deadline != nil {
		task.Deadline = updatedTask.Deadline
	}

	// Save updated task
	if err := config.DB.Save(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// ✅ Fix: DeleteTask now accepts `uuid.UUID`
func DeleteTask(id uuid.UUID) error {
	if err := config.DB.Delete(&models.Task{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
