package db

import (
	"fmt"

	"github.com/artyom-kalman/go-song-library/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration() error {
	logger.Logger.Debug("Running migration...")

	if databaseConnection == nil {
		return fmt.Errorf("Error running migration: database connection is closed")
	}

	databaseConnection.Exec("DROP TABLE IF EXISTS schema_migrations;")

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
	if err != nil {
		logger.Logger.Error(err.Error())
	} else {
		logger.Logger.Info("Migrated database schema")
	}

	return nil
}
