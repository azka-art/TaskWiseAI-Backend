package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Default cost for bcrypt
const bcryptCost = 12

// HashPassword mengenkripsi password dengan bcrypt
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	// Validate the password
	if err := ValidatePassword(password); err != nil {
		return "", err
	}

	// Generate bcrypt hash with appropriate cost factor
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPasswordHash membandingkan password dengan hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateSecureToken generates a cryptographically secure token for various purposes
// (like password reset, email verification, etc.)
func GenerateSecureToken() (string, error) {
	// TODO: Implement secure token generation, possibly using crypto/rand
	// For now, returning a placeholder
	return "", errors.New("not implemented")
}

// CompareHashAndPassword is a wrapper around bcrypt's CompareHashAndPassword
// that returns a more user-friendly error message
func CompareHashAndPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("incorrect password")
		}
		return err
	}
	return nil
}

// IsHashedPassword checks if a string is already a bcrypt hash
func IsHashedPassword(str string) bool {
	// bcrypt hashes start with $2a$, $2b$, or $2y$
	return len(str) == 60 && (str[:4] == "$2a$" || str[:4] == "$2b$" || str[:4] == "$2y$")
}
