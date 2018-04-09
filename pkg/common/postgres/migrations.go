package postgres

import (
	"github.com/build-tanker/tanker/pkg/common/config"
	_ "github.com/lib/pq" // postgres driver
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file" // get db migration from path

	"database/sql"
	"log"
)

const migrationsPath = "file://./pkg/migrations"

// RunDatabaseMigrations - run the next migration, needs to be run multiple times if there are multiple
func RunDatabaseMigrations(conf *config.Config) error {
	db, err := sql.Open("postgres", conf.ConnectionURL())

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		log.Println("Sadly, found no new migrations to run")
		return nil
	}

	if err != nil {
		return err
	}

	log.Println("Migration has been successfully done")
	return nil
}

// RollbackDatabaseMigration - rollback the database migration
func RollbackDatabaseMigration(conf *config.Config) error {
	m, err := migrate.New(migrationsPath, conf.ConnectionURL())
	if err != nil {
		return err
	}

	if err := m.Steps(-1); err != nil {
		log.Println("There was an error rolling back", err.Error())
		return nil
	}

	log.Println("Rollback Successful")
	return nil
}
