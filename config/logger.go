package config

import (
	"log"
	"os"
)

// InitLogger sets up the global logger
func InitLogger() {
	log.SetOutput(os.Stdout)            // Log to console
	log.SetFlags(log.Ldate | log.Ltime) // Format: YYYY/MM/DD HH:MM:SS
	log.Println("✅ Logger initialized.")
}
