package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/build-tanker/tanker/pkg/appcontext"
	"github.com/build-tanker/tanker/pkg/config"
	"github.com/build-tanker/tanker/pkg/logger"
	"github.com/build-tanker/tanker/pkg/postgres"
	"github.com/build-tanker/tanker/pkg/server"
)

func main() {
	config := config.NewConfig([]string{".", "..", "../.."})
	logger := logger.NewLogger(config, os.Stdout)
	ctx := appcontext.NewAppContext(config, logger)
	db := postgres.NewPostgres(logger, config.Database().ConnectionURL(), config.Database().MaxPoolSize())
	server := server.NewServer(ctx, db)

	logger.Infoln("Starting tanker")

	app := cli.NewApp()
	app.Name = "tanker"
	app.Version = "0.0.1"
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
