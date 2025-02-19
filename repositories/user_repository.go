package repositories

import (
	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
)

// CreateUser menyimpan user baru ke database
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// GetUserByEmail mencari user berdasarkan email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// GetAllUsers mengambil semua pengguna
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}
