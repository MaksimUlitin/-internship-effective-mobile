package config

import (
	"effectiveMobileTask/lib/logger"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
)

type Config struct {
	DB          DBConfig
	Server      ServerConfig
	ExternalAPI ExternalAPIConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	URL      string
}

type ServerConfig struct {
	Port           string
	MockServerPort string
}

type ExternalAPIConfig struct {
	BaseURL string
	InfoURL string
}

var AppConfig Config

func LoadConfigEnv() {
	if err := godotenv.Load(); err != nil {
		logger.Error("not found .env file", slog.Any("err", err))
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		DB: DBConfig{
			Host:     getEnvOrDefault("DB_HOST", ""),
			Port:     getEnvOrDefault("DB_PORT", ""),
			User:     getEnvOrDefault("DB_USER", ""),
			Password: getEnvOrDefault("DB_PASSWORD", ""),
			Name:     getEnvOrDefault("DB_NAME", ""),
			URL:      getEnvOrDefault("DATABASE_URL", ""),
		},
		Server: ServerConfig{
			Port:           getEnvOrDefault("SERVER_PORT", ""),
			MockServerPort: getEnvOrDefault("SERVER_MOCK_SERVER_PORT", ""),
		},
		ExternalAPI: ExternalAPIConfig{
			BaseURL: getEnvOrDefault("EXTERNAL_API_BASE_URL", ""),
			InfoURL: getEnvOrDefault("EXTERNAL_API_INFO_PATH", ""),
		},
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
