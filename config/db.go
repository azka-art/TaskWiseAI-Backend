package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection instance
var DB *gorm.DB

func ConnectDatabase() {
	// ✅ Force reload the latest `.env`
	err := godotenv.Overload()
	if err != nil {
		log.Println("⚠️ No .env file found. Using system environment variables.")
	} else {
		log.Println("✅ .env file loaded successfully!")
	}

	// Fetch environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// ✅ Trim any extra spaces and comments
	password = strings.TrimSpace(password)     // Remove spaces
	password = strings.Split(password, "#")[0] // Remove any inline comment
	password = strings.Trim(password, `"'`)    // Remove extra quotes

	// Debugging: Print all values
	log.Println("📌 Debug: Loaded Environment Variables")
	log.Printf("DB_HOST: %s", host)
	log.Printf("DB_USER: %s", user)
	log.Printf("DB_PASSWORD: %s", password) // Only for debugging, remove later!
	log.Printf("DB_NAME: %s", dbname)
	log.Printf("DB_PORT: %s", port)
	log.Printf("DB_SSLMODE: %s", sslmode)

	// If any required variable is missing, stop execution
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatalf("❌ Missing database configuration! Check your .env file.")
	}

	// ✅ Correctly formatted DSN string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	// Print DSN before connecting
	log.Printf("📌 DSN: %s", dsn)

	// Open the database connection using GORM
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}

	// Verify connection
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	// ✅ Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	// Assign database connection to global variable
	DB = database
	log.Println("✅ Database connection established successfully! 🚀")
}
