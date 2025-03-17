package db

import (
	"context"
	"database/sql"
	"time"
)

func (db *DatabaseConnection) Query(q string) (*sql.Rows, error) {
	c, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := db.conn.PingContext(c)
	if err != nil {
		return nil, err
	}

	rows, err := db.conn.QueryContext(c, q)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (db *DatabaseConnection) QueryContext(c context.Context, q string) (*sql.Rows, error) {
	err := db.conn.PingContext(c)
	if err != nil {
		return nil, err
	}

	rows, err := db.conn.QueryContext(c, q)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
