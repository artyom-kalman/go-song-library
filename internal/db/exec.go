package db

func (db *DatabaseConnection) Exec(query string) (int, error) {
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
