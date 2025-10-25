package internal

import (
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Port      string
	AppDomain string
	DSN       string
	JWKSURL   string
}

func NewConfig() *Config {
	// Load .env file in development environment
	// For production like ECS/EKS, environment variables will be set directly
	_ = godotenv.Load()

	return &Config{
		Port:      getEnv("PORT", "8080"),
		AppDomain: getEnvRequired("APP_DOMAIN"),
		DSN:       getEnvRequired("DSN"),
		JWKSURL:   getEnvRequired("JWKS_URL"),
	}
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Required environment variable not set: " + key)
	}
	return value
}
