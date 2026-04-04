package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://localhost:5432/vibeapp?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "secret"),
	}
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}