package postgres

import (
	_ "github.com/lib/pq" // postgres driver
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file" // get db migration from path

	"database/sql"

	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
)

const migrationsPath = "file://./db/migrations"

// RunDatabaseMigrations : run each migration on the db
func RunDatabaseMigrations() error {
	db, err := sql.Open("postgres", config.Database().ConnectionURL())

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		logger.Info("Sadly, found no new migrations to run")
		return nil
	}

	if err != nil {
		return err
	}

	logger.Info("Migration has been successfully done")
	return nil
}

// RollbackDatabaseMigration : rollback the latest migration
func RollbackDatabaseMigration() error {
	m, err := migrate.New(migrationsPath, config.Database().ConnectionURL())
	if err != nil {
		return err
	}

	if err := m.Steps(-1); err != nil {
		logger.Info("We have already removed all migrations")
		return nil
	}

	logger.Info("Rollback Successful")
	return nil
}
