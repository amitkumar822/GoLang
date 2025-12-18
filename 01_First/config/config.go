package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	AppPort       string
	MongoURI      string
	MongoDB       string
	JWTSecret     string
	JWTExpireHours int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists (ignore error if file doesn't exist)
	_ = godotenv.Load()

	// Get JWT_EXPIRE_HOURS and convert to int, default to 24
	jwtExpireHours := 24
	if hours := os.Getenv("JWT_EXPIRE_HOURS"); hours != "" {
		if parsed, err := strconv.Atoi(hours); err == nil {
			jwtExpireHours = parsed
		}
	}

	config := &Config{
		AppPort:        getEnv("APP_PORT", "8080"),
		MongoURI:       getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:        getEnv("MONGO_DB", "userdb"),
		JWTSecret:      getEnv("JWT_SECRET", "supersecretkey"),
		JWTExpireHours: jwtExpireHours,
	}

	// Validate required fields
	if config.JWTSecret == "supersecretkey" {
		log.Println("WARNING: Using default JWT_SECRET. Please set a secure secret in production!")
	}

	return config
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

