package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from .env
func LoadConfig() {
	_ = godotenv.Load(".env") // Optional, defaults to environment vars
}

// GetConfig retrieves a value from environment variables
func GetConfig(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
