package controllers

import (
	"e-learning/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeMessage provides a basic API response and tests the database connection.
func WelcomeMessage(c *gin.Context) {
	// Test the database connection
	err := database.DB.Exec("SELECT 1").Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database connection failed", "error": err.Error()})
		return
	}

	// Respond with a welcome message
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the E-Learning API!",
		"status":  "Database connected successfully",
	})
}
