package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration() error {
	if databaseConnection == nil {
		return fmt.Errorf("Error running migration: database connection is closed")
	}

	driver, err := postgres.WithInstance(databaseConnection.connection, &postgres.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations/",
		"song_lib",
		driver,
	)
	if err != nil {
		return err
	}

	err = migration.Up()

	return nil
}
