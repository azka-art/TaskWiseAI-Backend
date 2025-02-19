package services

import (
	"errors"
	"time"

	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/models"
	"github.com/azka-art/taskwise-backend/utils"
	"github.com/golang-jwt/jwt/v5"
)

// RegisterUser mendaftarkan pengguna baru
func RegisterUser(user models.User) (models.User, error) {
	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword

	// Simpan ke database
	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// LoginUser memeriksa kredensial pengguna dan mengembalikan token JWT
func LoginUser(credentials models.LoginRequest) (string, error) {
	var user models.User

	// Cari pengguna berdasarkan email
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	// Verifikasi password
	if !utils.CheckPassword(credentials.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := "your_secret_key"
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
