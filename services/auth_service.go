package services

import (
	"errors"
	"os"
	"time"

	"github.com/azka-art/taskwise-backend/models"
	"github.com/azka-art/taskwise-backend/repositories"
	"github.com/azka-art/taskwise-backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// RegisterUser registers a new user with validation
func RegisterUser(user models.User) (models.User, error) {
	// Validate email
	if err := utils.ValidateEmail(user.Email); err != nil {
		return models.User{}, err
	}

	// Validate username
	if err := utils.ValidateUsername(user.Username); err != nil {
		return models.User{}, err
	}

	// Validate password
	if err := utils.ValidatePassword(user.Password); err != nil {
		return models.User{}, err
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = models.RoleMember
	}

	// Generate UUID if not provided
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	// Check if email already exists
	existingUser, _ := repositories.GetUserByEmail(user.Email)
	if existingUser != nil {
		return models.User{}, errors.New("email already in use")
	}

	// Hash password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword

	// Save to database
	if err := repositories.CreateUser(&user); err != nil {
		return models.User{}, err
	}

	// Clear password before returning
	user.Password = ""
	return user, nil
}

// LoginUser verifies user credentials and returns a JWT token
func LoginUser(credentials models.LoginRequest) (string, models.User, error) {
	// Validate email format
	if err := utils.ValidateEmail(credentials.Email); err != nil {
		return "", models.User{}, err
	}

	user, err := repositories.GetUserByEmail(credentials.Email)
	if err != nil || user == nil {
		return "", models.User{}, errors.New("invalid email or password")
	}

	// Verify password (Changed `CheckPassword` → `CheckPasswordHash`)
	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		return "", models.User{}, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := generateToken(user)
	if err != nil {
		return "", models.User{}, err
	}

	// Clear password before returning
	user.Password = ""
	return token, *user, nil
}

// generateToken creates a JWT for a given user
func generateToken(user *models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your_secret_key"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	})

	return token.SignedString([]byte(jwtSecret))
}

// ValidateToken verifies a JWT token and returns claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your_secret_key"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ChangePassword allows users to update their password
func ChangePassword(userID uuid.UUID, oldPassword, newPassword string) error {
	// Validate new password
	if err := utils.ValidatePassword(newPassword); err != nil {
		return err
	}

	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Verify old password (Changed `CheckPassword` → `CheckPasswordHash`)
	if !utils.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("incorrect current password")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update password
	user.Password = hashedPassword
	return repositories.UpdateUser(user)
}
