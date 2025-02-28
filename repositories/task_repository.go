package repositories

import (
	"errors"

	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateTask menambahkan tugas baru ke database
func CreateTask(task *models.Task) error {
	// Set default values if needed
	if task.Priority == "" {
		task.Priority = models.PriorityMedium
	}

	if task.Status == "" {
		task.Status = models.StatusPending
	}

	// Generate UUID if not provided
	if task.ID == uuid.Nil {
		task.ID = uuid.New()
	}

	// Validate task before saving
	if err := validateTask(task); err != nil {
		return err
	}

	return config.DB.Create(task).Error
}

// GetAllTasks mengambil semua tugas
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Preload("User").Find(&tasks).Error
	return tasks, err
}

// GetTasksByUserID mengambil tugas berdasarkan user ID
func GetTasksByUserID(userID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("created_by = ?", userID).Preload("User").Find(&tasks).Error
	return tasks, err
}

// GetTasksByStatus mengambil tugas berdasarkan status
func GetTasksByStatus(status models.Status) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("status = ?", status).Preload("User").Find(&tasks).Error
	return tasks, err
}

// GetTasksByPriority mengambil tugas berdasarkan prioritas
func GetTasksByPriority(priority models.Priority) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("priority = ?", priority).Preload("User").Find(&tasks).Error
	return tasks, err
}

// GetTaskWithComments mengambil tugas dengan komentar
func GetTaskWithComments(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := config.DB.Preload("Comments").Preload("Comments.User").Preload("User").Where("id = ?", id).First(&task).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

// GetTaskByID mencari tugas berdasarkan ID
func GetTaskByID(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := config.DB.Preload("User").Where("id = ?", id).First(&task).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

// UpdateTask memperbarui tugas berdasarkan ID
func UpdateTask(task *models.Task) error {
	// Check if task exists
	var count int64
	if err := config.DB.Model(&models.Task{}).Where("id = ?", task.ID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New("task not found")
	}

	// Validate task before updating
	if err := validateTask(task); err != nil {
		return err
	}

	// Update only specific fields to prevent overwriting data that shouldn't be changed
	return config.DB.Model(&models.Task{}).Where("id = ?", task.ID).Updates(map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"priority":    task.Priority,
		"status":      task.Status,
		"deadline":    task.Deadline,
	}).Error
}

// UpdateTaskStatus memperbarui status tugas
func UpdateTaskStatus(id uuid.UUID, status models.Status) error {
	// Validate status
	if status != models.StatusPending && status != models.StatusInProgress && status != models.StatusDone {
		return errors.New("invalid status value")
	}

	result := config.DB.Model(&models.Task{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

// DeleteTask menghapus tugas berdasarkan ID
func DeleteTask(id uuid.UUID) error {
	// Begin a transaction to delete the task and its comments
	tx := config.DB.Begin()

	// Delete all comments associated with the task
	if err := tx.Where("task_id = ?", id).Delete(&models.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete the task
	result := tx.Where("id = ?", id).Delete(&models.Task{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("task not found")
	}

	return tx.Commit().Error
}

// GetTaskCount menghitung jumlah total tugas
func GetTaskCount() (int64, error) {
	var count int64
	err := config.DB.Model(&models.Task{}).Count(&count).Error
	return count, err
}

// GetPaginatedTasks mengambil tugas dengan pagination
func GetPaginatedTasks(page, size int) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	// Count total records
	if err := config.DB.Model(&models.Task{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * size

	// Get paginated data
	err := config.DB.Preload("User").
		Offset(offset).
		Limit(size).
		Order("created_at DESC").
		Find(&tasks).Error

	return tasks, total, err
}

// validateTask validates a task
func validateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}

	if task.CreatedBy == uuid.Nil {
		return errors.New("creator is required")
	}

	// Check priority value
	if task.Priority != "" &&
		task.Priority != models.PriorityLow &&
		task.Priority != models.PriorityMedium &&
		task.Priority != models.PriorityHigh {
		return errors.New("invalid priority value")
	}

	// Check status value
	if task.Status != "" &&
		task.Status != models.StatusPending &&
		task.Status != models.StatusInProgress &&
		task.Status != models.StatusDone {
		return errors.New("invalid status value")
	}

	return nil
}
