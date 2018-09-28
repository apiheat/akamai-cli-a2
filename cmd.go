package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

func cmdReport(c *cli.Context) error {
	return reportProperty(c)
}

func cmdReset(c *cli.Context) error {
	return resetProperty(c)
}

func reportRespParse(in string) (data reportResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func reportProperty(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/%s", URL, common.SetIntID(c, "Please provide ID for property"))

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := reportRespParse(data)
	common.ErrorCheck(err)

	common.OutputJSON(result)

	return nil
}

func resetProperty(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/%s/reset", URL, common.SetIntID(c, "Please provide ID for property"))

	if debug {
		println(urlStr)
	}

	_, respCode := fetchData(urlStr, "POST", nil)

	if debug {
		println(respCode)
	}

	if respCode == 204 {
		log.Println("Reset request was completed successfully")
	} else {
		log.Printf("Something went wrong! Response code is: %v", respCode)
	}

	return nil
}

func fetchData(urlPath, method string, body io.Reader) (string, int) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, body)
	common.ErrorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	common.ErrorCheck(err)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt), resp.StatusCode
}
