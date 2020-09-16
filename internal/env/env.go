package env

import (
	"os"
	"strconv"
)

type ConfigAPI struct {
	TokenApi string
	ChatId   int64
}

// New returns a new Config struct
func New() ConfigAPI {
	return ConfigAPI{
		TokenApi: getEnv("TOKENAPI", ""),
		ChatId:   getEnvAsInt64("CHATID", 0),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt64(name string, defaultVal int) int64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return int64(value)
	}

	return int64(defaultVal)
}
