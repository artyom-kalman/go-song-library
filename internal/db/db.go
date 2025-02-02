package db

import (
	"database/sql"
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/config"

	_ "github.com/lib/pq"
)

var openedConn *DBConnection

type DBConnection struct {
	connection *sql.DB
}

func ConnectToDB() (*DBConnection, error) {
	if openedConn != nil {
		return openedConn, nil
	}

	config, err := config.LoadDBConfig()
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Name,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	openedConn = &DBConnection{
		connection: db,
	}

	return openedConn, nil
}

func (db *DBConnection) Query(query string, dest ...any) (*sql.Rows, error) {
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
