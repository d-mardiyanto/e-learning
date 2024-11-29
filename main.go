package main

import (
	"e-learning/config"
	"e-learning/database"
	"e-learning/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	database.ConnectDB()

	// Set up Gin
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	port := config.GetConfig("PORT", "8080")
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
