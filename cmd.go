package main

import (
	common "github.com/apiheat/akamai-cli-common"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

func cmdReport(c *cli.Context) error {
	return reportProperty(c)
}

func cmdReset(c *cli.Context) error {
	return resetProperty(c)
}

func reportProperty(c *cli.Context) error {
	report, err := apiClient.A2.ReportProperty(common.SetIntID(c, "Please provide ID for property"))
	common.ErrorCheck(err)

	common.PrintJSON(report.Body)

	return nil
}

func resetProperty(c *cli.Context) error {
	report, err := apiClient.A2.ResetProperty(common.SetIntID(c, "Please provide ID for property"))
	common.ErrorCheck(err)

	if report.Response.StatusCode == 204 {
		log.Info("Reset request was completed successfully")
	} else {
		log.Info("Something went wrong! Response code is: %v", report.Response.StatusCode)
	}

	return nil
}
