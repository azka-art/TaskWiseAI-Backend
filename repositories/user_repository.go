package repositories

import (
	"errors"

	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
	"github.com/azka-art/taskwise-backend/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateUser menyimpan user baru ke database
func CreateUser(user *models.User) error {
	// Set default values
	if user.Role == "" {
		user.Role = models.RoleMember
	}

	// Generate UUID if not provided
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	// Validate user data
	if err := validateUser(user); err != nil {
		return err
	}

	// Check if email already exists
	existingUser, err := GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already in use")
	}

	// Check if username already exists
	existingUser, err = GetUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already taken")
	}

	// Hash password before storing
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return config.DB.Create(user).Error
}

// GetUserByID mencari user berdasarkan ID
func GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := config.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil when no record is found
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail mencari user berdasarkan email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil when no record is found
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername mencari user berdasarkan username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil when no record is found
		}
		return nil, err
	}

	return &user, nil
}

// GetAllUsers mengambil semua pengguna
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

// GetPaginatedUsers mengambil pengguna dengan pagination
func GetPaginatedUsers(page, size int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Count total records
	if err := config.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * size

	// Get paginated data
	err := config.DB.
		Offset(offset).
		Limit(size).
		Order("created_at DESC").
		Find(&users).Error

	return users, total, err
}

// UpdateUser memperbarui data pengguna
func UpdateUser(user *models.User) error {
	// Check if user exists
	existingUser, err := GetUserByID(user.ID)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	// Validate user data
	if err := validateUser(user); err != nil {
		return err
	}

	// Check if email is unique (if changed)
	if existingUser.Email != user.Email {
		emailUser, err := GetUserByEmail(user.Email)
		if err == nil && emailUser != nil && emailUser.ID != user.ID {
			return errors.New("email already in use")
		}
	}

	// Check if username is unique (if changed)
	if existingUser.Username != user.Username {
		usernameUser, err := GetUserByUsername(user.Username)
		if err == nil && usernameUser != nil && usernameUser.ID != user.ID {
			return errors.New("username already taken")
		}
	}

	// Update only these fields, not password or role
	return config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
	}).Error
}

// UpdateUserRole memperbarui role pengguna (admin only)
func UpdateUserRole(userID uuid.UUID, role models.UserRole) error {
	// Validate role
	if role != models.RoleAdmin && role != models.RoleMember {
		return errors.New("invalid role value")
	}

	// Check if user exists
	existingUser, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	return config.DB.Model(&models.User{}).Where("id = ?", userID).Update("role", role).Error
}

// UpdatePassword memperbarui password pengguna
func UpdatePassword(userID uuid.UUID, currentPassword, newPassword string) error {
	// Get user
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	// Verify current password
	if !utils.CheckPasswordHash(currentPassword, user.Password) {
		return errors.New("current password is incorrect")
	}

	// Validate new password
	if len(newPassword) < 6 {
		return errors.New("new password must be at least 6 characters")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update password
	return config.DB.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}

// DeleteUser menghapus pengguna berdasarkan ID
func DeleteUser(id uuid.UUID) error {
	// Check if user exists
	existingUser, err := GetUserByID(id)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	// Begin a transaction
	tx := config.DB.Begin()

	// Delete user's comments
	if err := tx.Where("user_id = ?", id).Delete(&models.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete user's tasks and their comments
	var tasks []models.Task
	if err := tx.Where("created_by = ?", id).Find(&tasks).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, task := range tasks {
		// Delete comments for this task
		if err := tx.Where("task_id = ?", task.ID).Delete(&models.Comment{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Delete the task
		if err := tx.Delete(&task).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Delete the user
	if err := tx.Delete(existingUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetUserCount returns the total number of users
func GetUserCount() (int64, error) {
	var count int64
	err := config.DB.Model(&models.User{}).Count(&count).Error
	return count, err
}

// validateUser validates user data
func validateUser(user *models.User) error {
	if user.Username == "" {
		return errors.New("username is required")
	}

	if len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	// Basic email format validation
	if !utils.IsValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	// Only validate password if it's provided (for create or password updates)
	if user.Password != "" && len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}
