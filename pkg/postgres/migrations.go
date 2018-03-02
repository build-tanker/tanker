package postgres

import (
	_ "github.com/lib/pq" // postgres driver
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file" // get db migration from path

	"database/sql"

	"github.com/build-tanker/tanker/pkg/appcontext"
)

const migrationsPath = "file://./pkg/postgres/migrations"

func RunDatabaseMigrations(ctx *appcontext.AppContext) error {
	db, err := sql.Open("postgres", ctx.GetConfig().Database().ConnectionURL())

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		ctx.GetLogger().Infoln("Sadly, found no new migrations to run")
		return nil
	}

	if err != nil {
		return err
	}

	ctx.GetLogger().Infoln("Migration has been successfully done")
	return nil
}

func RollbackDatabaseMigration(ctx *appcontext.AppContext) error {
	m, err := migrate.New(migrationsPath, ctx.GetConfig().Database().ConnectionURL())
	if err != nil {
		return err
	}

	if err := m.Steps(-1); err != nil {
		ctx.GetLogger().Infoln("We have already removed all migrations")
		return nil
	}

	ctx.GetLogger().Infoln("Rollback Successful")
	return nil
}
