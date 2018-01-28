package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
	"github.com/sudhanshuraheja/tanker/pkg/postgres"
	"github.com/sudhanshuraheja/tanker/pkg/server"
)

func main() {
	config := config.NewConfig()
	logger := logger.NewLogger(config)
	ctx := appcontext.NewAppContext(config, logger)
	db := postgres.NewPostgres(*logger, config.Database().ConnectionURL(), config.Database().MaxPoolSize())
	server := server.NewServer(ctx, db)

	logger.Infoln("Starting sample-cli")

	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this service saves files and makes them available for distribution"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return server.StartAPIServer()
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration()
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

