package config

import (
	"os"

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

	return nil
}

func GetDBConfig() (*DBConfig, error) {
	if !isConfigLoaded {
		err := LoadConfig()
		if err != nil {
			return nil, err
		}
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}, nil
}

func GetServerConfig() (*ServerConfig, error) {
	if !isConfigLoaded {
		if err := LoadConfig(); err != nil {
			return nil, err
		}
	}

	return &ServerConfig{
		Port: os.Getenv("PORT"),
	}, nil
}
