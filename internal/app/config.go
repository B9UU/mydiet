package app

import (
	"os"
	"strconv"
	"time"
)

// Config holds application configuration
type Config struct {
	Database DatabaseConfig
	Logging  LoggingConfig
	UI       UIConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Path            string
	ConnectionRetries int
	ConnectionTimeout time.Duration
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level    string
	FilePath string
	Console  bool
}

// UIConfig holds UI configuration
type UIConfig struct {
	DefaultView string
	HelpEnabled bool
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Path:              getEnvOrDefault("DB_PATH", "./nutrition.db"),
			ConnectionRetries: getEnvAsIntOrDefault("DB_CONNECTION_RETRIES", 3),
			ConnectionTimeout: time.Duration(getEnvAsIntOrDefault("DB_CONNECTION_TIMEOUT_SECONDS", 5)) * time.Second,
		},
		Logging: LoggingConfig{
			Level:    getEnvOrDefault("LOG_LEVEL", "info"),
			FilePath: getEnvOrDefault("LOG_FILE", "debug.log"),
			Console:  getEnvAsBoolOrDefault("LOG_CONSOLE", true),
		},
		UI: UIConfig{
			DefaultView: getEnvOrDefault("DEFAULT_VIEW", "details"),
			HelpEnabled: getEnvAsBoolOrDefault("HELP_ENABLED", true),
		},
	}
}

// getEnvOrDefault gets an environment variable or returns a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsIntOrDefault gets an environment variable as int or returns default
func getEnvAsIntOrDefault(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// getEnvAsBoolOrDefault gets an environment variable as bool or returns default
func getEnvAsBoolOrDefault(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}