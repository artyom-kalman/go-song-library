package db

import "database/sql"

func (db *DatabaseConnection) Query(query string) (*sql.Rows, error) {
	err := db.conn.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
