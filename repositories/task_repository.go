package repositories

import (
	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
)

// CreateTask menambahkan tugas baru ke database
func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

// GetAllTasks mengambil semua tugas
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Find(&tasks).Error
	return tasks, err
}

// GetTaskByID mencari tugas berdasarkan ID
func GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	err := config.DB.First(&task, id).Error
	return &task, err
}

// UpdateTask memperbarui tugas berdasarkan ID
func UpdateTask(task *models.Task) error {
	return config.DB.Save(task).Error
}

// DeleteTask menghapus tugas berdasarkan ID
func DeleteTask(id uint) error {
	return config.DB.Delete(&models.Task{}, id).Error
}
