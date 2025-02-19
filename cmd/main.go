package main

import (
	"log"

	"github.com/azka-art/taskwise-backend/config"
	"github.com/azka-art/taskwise-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Initialize Gin router
	r := gin.Default()

	// Register all routes
	routes.SetupRoutes(r)

	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(r.Run(":8080")) // Ensures server exits on failure
}
