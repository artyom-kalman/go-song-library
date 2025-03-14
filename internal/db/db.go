package db

import (
	"database/sql"

	"github.com/artyom-kalman/go-song-library/internal/config"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	conn *sql.DB
}

var databaseConnection *DatabaseConnection

func Database() *DatabaseConnection {
	if databaseConnection == nil {
		databaseConfig, _ := config.GetDBConfig()
		ConnectToDatabase(databaseConfig)
	}

	return databaseConnection
}
