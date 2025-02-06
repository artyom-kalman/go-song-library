package config

import (
	"fmt"
	"os"

	"github.com/artyom-kalman/go-song-library/pkg/logger"
	"github.com/joho/godotenv"
)

var isConfigLoaded bool = false

func LoadConfig() error {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			return err
		}
		isConfigLoaded = true
	}

	if _, err := os.Stat("../../.env"); err == nil {
		err := godotenv.Load("../../.env")
		if err != nil {
			return err
		}
		isConfigLoaded = true
	}

	if !isConfigLoaded {
		return fmt.Errorf("Failed to load .env")
	}

	logger.Logger.Info("Loaded config from .env file")

	return nil
}

func GetDBConfig() (*DBConfig, error) {
	if !isConfigLoaded {
		err := LoadConfig()
		if err != nil {
			return nil, err
		}
	}

	logger.Logger.Info("Loaded database configuration")
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("POSTGRES_DB"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}, nil
}

func GetServerConfig() (*ServerConfig, error) {
	if !isConfigLoaded {
		if err := LoadConfig(); err != nil {
			return nil, err
		}
	}

	logger.Logger.Info("Loaded server configuration")
	return &ServerConfig{
		Port: os.Getenv("APP_PORT"),
	}, nil
}

func GetOpenAPIURL() string {
	return os.Getenv("OPENAPI_URL")
}
