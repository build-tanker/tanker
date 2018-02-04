package main

import (
	"os"
	"sort"

	"github.com/urfave/cli"

	"source.golabs.io/core/tanker/pkg/appcontext"
	"source.golabs.io/core/tanker/pkg/config"
	"source.golabs.io/core/tanker/pkg/logger"
	"source.golabs.io/core/tanker/pkg/uploader"
)

func main() {
	config := config.NewConfig()
	logger := logger.NewLogger(config)
	ctx := appcontext.NewAppContext(config, logger)

	logger.Infoln("Starting shipper")

	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this binary uploads builds for distribution"

	app.Action = func(c *cli.Context) error {
		logger.Infoln("Getting ready to ship")
		err := uploader.Upload(c.String("key"), c.String("bundle"), c.String("file"))
		if err != nil {
			logger.Infoln(err)
		}
		return nil
	}

	app.Commands = []cli.Command{}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "key, k",
			Usage: "access key for authentication",
		},
		cli.StringFlag{
			Name:  "file, f",
			Usage: "file to be uploaded",
		},
		cli.StringFlag{
			Name:  "bundle, b",
			Usage: "app bundle to link to",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
