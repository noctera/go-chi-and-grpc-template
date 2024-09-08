package config

import (
	"log"
	"os"
)

// Config holds the configuration values
type Config struct {
	DatabaseURL string
}

// LoadConfig reads configuration from environment variables
func LoadConfig() Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	return Config{
		DatabaseURL: dbURL,
	}
}
