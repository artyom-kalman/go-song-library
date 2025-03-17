package db

import (
	"context"
	"time"
)

func (db *DatabaseConnection) Exec(q string) (int, error) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := db.conn.PingContext(c)
	if err != nil {
		return 0, err
	}

	result, err := db.conn.ExecContext(c, q)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected), nil
}

func (db *DatabaseConnection) ExecContext(c context.Context, q string) (int, error) {
	err := db.conn.PingContext(c)
	if err != nil {
		return 0, err
	}

	result, err := db.conn.ExecContext(c, q)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected), nil
}
