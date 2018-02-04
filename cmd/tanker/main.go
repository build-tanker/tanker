package main

import (
	"os"

	"github.com/urfave/cli"

	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/config"
	"source.golabs.io/core/tanker/pkg/logger"
	"source.golabs.io/core/tanker/pkg/postgres"
	"source.golabs.io/core/tanker/pkg/server"
)

func main() {
	config := config.NewConfig()
	logger := logger.NewLogger(config)
	ctx := appcontext.NewAppContext(config, logger)
	db := postgres.NewPostgres(logger, config.Database().ConnectionURL(), config.Database().MaxPoolSize())
	server := server.NewServer(ctx, db)

	logger.Infoln("Starting tanker")

	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this service saves files and makes them available for distribution"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations(ctx)
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration(ctx)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
