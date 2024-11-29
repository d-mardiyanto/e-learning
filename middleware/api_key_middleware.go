package middleware

import (
	"e-learning/database"
	"e-learning/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIKeyMiddleware validates requests based on an API key in the headers.
func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the API key from the request header
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key is required"})
			c.Abort()
			return
		}

		// Check the database for the API key
		var key models.APIKey
		if err := database.DB.Where("token = ?", apiKey).First(&key).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		// API key is valid, continue to the next middleware/handler
		c.Next()
	}
}
