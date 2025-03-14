package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/artyom-kalman/go-song-library/internal/config"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

var (
	maxRetries = 5
	retryDelay = 2 * time.Second
)

func ConnectToDatabase(config *config.DBConfig) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Name,
	)

	logger.Info("Attempt to connect to database...")

	var db *sql.DB
	var err error
	for range maxRetries {
		db, err = sql.Open("postgres", connectionString)
		if err = db.Ping(); err != nil {
			logger.Info("Failed to connect: %v. Retrying...", err)
			time.Sleep(retryDelay)
			continue
		}
		databaseConnection = &DatabaseConnection{db}
		break
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	logger.Info("Initialized database connection")

	return nil
}
