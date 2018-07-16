package main

import (
	"os"
	"sort"
	"time"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var (
	colorOn, raw, debug       bool
	version, appName          string
	configSection, configFile string
	edgeConfig                edgegrid.Config
)

// Constants
const (
	URL     = "/adaptive-acceleration/v1/properties"
	padding = 3
)

type reportResponse struct {
	CreationDate           time.Time `json:"creationDate"`
	IsActive               bool      `json:"isActive"`
	LastModifiedDate       time.Time `json:"lastModifiedDate"`
	LastReset              time.Time `json:"lastReset"`
	ZoneDeployDate         time.Time `json:"zoneDeployDate"`
	Version                int       `json:"version"`
	CommonPreconnectHeader []string  `json:"commonPreconnectHeader"`
	CommonPushedResources  []string  `json:"commonPushedResources"`
	PageSpecificRules      []struct {
		BasePageURL                  string   `json:"basePageURL"`
		PageSpecificPreconnectHeader []string `json:"pageSpecificPreconnectHeader"`
		PageSpecificPushedResources  []string `json:"pageSpecificPushedResources"`
	} `json:"pageSpecificRules"`
}

func main() {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	appName = "akamai-a2"
	if inCLI {
		appName = "akamai a2"
	}

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Usage = "A CLI to interact with Akamai Adaptive Acceleration"
	app.Version = version
	app.Copyright = ""
	app.Authors = []cli.Author{
		{
			Name: "Petr Artamonov",
		},
		{
			Name: "Rafal Pieniazek",
		},
	}

	dir, _ := homedir.Dir()
	dir += string(os.PathSeparator) + ".edgerc"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &configSection,
			EnvVar:      "AKAMAI_EDGERC_SECTION",
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       dir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &configFile,
			EnvVar:      "AKAMAI_EDGERC",
		},
		cli.BoolFlag{
			Name:        "no-color",
			Usage:       "Disable color output",
			Destination: &colorOn,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Debug info",
			Destination: &debug,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "report",
			Aliases: []string{"r"},
			Usage:   "Get a report for property [ID]",
			Action:  cmdReport,
		},
		{
			Name:    "reset",
			Aliases: []string{"rm"},
			Usage:   "Reset all existing info for property [ID]",
			Action:  cmdReset,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		if c.Bool("no-color") {
			color.NoColor = true
		}

		config(configFile, configSection)
		return nil
	}

	app.Run(os.Args)
}
