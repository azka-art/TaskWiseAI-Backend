package utils

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

// ValidateEmail memastikan format email valid
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	email = strings.TrimSpace(email)
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// IsValidEmail checks if the email has a valid format (returns boolean)
func IsValidEmail(email string) bool {
	return ValidateEmail(email) == nil
}

// ValidatePassword validates a password with stricter rules
func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Uncomment for stronger password validation (if desired)

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// ValidateUsername validates a username
func ValidateUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}

	username = strings.TrimSpace(username)
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}

	if len(username) > 50 {
		return errors.New("username cannot exceed 50 characters")
	}

	// Check if username has valid characters (alphanumeric, underscores, hyphens)
	usernameRegex := `^[a-zA-Z0-9_-]+$`
	re := regexp.MustCompile(usernameRegex)
	if !re.MatchString(username) {
		return errors.New("username can only contain letters, numbers, underscores, and hyphens")
	}

	return nil
}

// ValidateUUID validates if a string is a valid UUID
func ValidateUUID(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid UUID format")
	}

	return nil
}

// StringToUUID converts a string to UUID and validates it
func StringToUUID(id string) (uuid.UUID, error) {
	if err := ValidateUUID(id); err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(id)
}

// ValidateURL validates if a string is a valid URL
func ValidateURL(rawURL string) error {
	if rawURL == "" {
		return errors.New("URL cannot be empty")
	}

	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return errors.New("invalid URL format")
	}

	return nil
}

// ValidateNotEmpty validates if a string is not empty
func ValidateNotEmpty(fieldName, value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New(fieldName + " cannot be empty")
	}
	return nil
}

// HasMinLength validates if a string meets the minimum length
func HasMinLength(fieldName, value string, minLength int) error {
	if len(value) < minLength {
		return errors.New(fieldName + " must be at least " + string(rune(minLength)) + " characters long")
	}
	return nil
}

// HasMaxLength validates if a string doesn't exceed the maximum length
func HasMaxLength(fieldName, value string, maxLength int) error {
	if len(value) > maxLength {
		return errors.New(fieldName + " cannot exceed " + string(rune(maxLength)) + " characters")
	}
	return nil
}
