package db

func (db *DatabaseConnection) Exec(query string) (int, error) {
	err := db.conn.Ping()
	if err != nil {
		return 0, err
	}

	result, err := db.conn.Exec(query)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected), nil
}
