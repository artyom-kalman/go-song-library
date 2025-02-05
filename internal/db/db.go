package db

import (
	"database/sql"
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/config"
	"github.com/artyom-kalman/go-song-library/pkg/logger"

	_ "github.com/lib/pq"
)

var databaseConnection *DBConnection

type DBConnection struct {
	connection *sql.DB
}

func GetDatabase() *DBConnection {
	if databaseConnection == nil {
		databaseConfig, _ := config.GetDBConfig()
		InitDatabase(databaseConfig)
	}

	return databaseConnection
}

func InitDatabase(config *config.DBConfig) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Name,
	)

	logger.Debug("Attempt to connect to database...")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	databaseConnection = &DBConnection{
		connection: db,
	}

	logger.Info("Initialized database connection")

	return nil
}

func (db *DBConnection) Query(query string) (*sql.Rows, error) {
	err := db.connection.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.connection.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (db *DBConnection) Exec(query string) (int, error) {
	err := db.connection.Ping()
	if err != nil {
		return 0, err
	}

	result, err := db.connection.Exec(query)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected), nil
}
