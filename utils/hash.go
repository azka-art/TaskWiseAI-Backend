package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword mengenkripsi password menggunakan bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword membandingkan password yang diinput dengan yang telah di-hash
func CheckPassword(inputPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
