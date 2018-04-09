package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/build-tanker/tanker/pkg/common/config"
	"github.com/build-tanker/tanker/pkg/common/postgres"
	"github.com/build-tanker/tanker/pkg/common/server"
)

func main() {
	cnf := config.New([]string{".", "..", "../.."})
	db := postgres.New(cnf.ConnectionString(), cnf.MaxPoolSize())
	server := server.New(cnf, db)

	fmt.Println("Starting tanker")

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
				return postgres.RunDatabaseMigrations(cnf)
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration(cnf)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
