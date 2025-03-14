package db

import "database/sql"

func (db *DatabaseConnection) Query(query string) (*sql.Rows, error) {
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
