package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// AppConfig holds the configuration values for the application
type AppConfig struct {
	PORT string
}

// Global variable to store the application configuration
var Config AppConfig

// LoadConfig initializes the configuration by loading environment variables
func LoadConfig() error {
	// Load the .env file into the system environment
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file")
	}

	// Initialize Config with environment variables
	Config = AppConfig{
		PORT: getEnv("PORT", "6060"),
	}

	return nil
}

// getEnv helps retrieve environment variables and provides a default if not found
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
