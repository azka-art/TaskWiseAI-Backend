package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
	err := godotenv.Overload()
	if err != nil {
		log.Println("⚠️ No .env file found. Using system environment variables.")
	} else {
		log.Println("✅ .env file loaded successfully!")
	}
}
