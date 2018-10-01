package main

import (
	"os"
	"sort"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var (
	apiClient       *edgegrid.Client
	appName, appVer string
)

func main() {
	app := common.CreateNewApp(appName, "A CLI to interact with Akamai Adaptive Acceleration", appVer)
	app.Flags = common.CreateFlags()

	app.Commands = []cli.Command{
		{
			Name:    "report",
			Aliases: []string{"r"},
			Usage:   "Get a report for property [ID]",
			Action:  cmdReport,
		},
		{
			Name:    "reset",
			Aliases: []string{"re"},
			Usage:   "Reset all existing info for property [ID]",
			Action:  cmdReset,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		var err error

		apiClient, err = common.EdgeClientInit(c.GlobalString("config"), c.GlobalString("section"), c.GlobalString("debug"))

		if err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
