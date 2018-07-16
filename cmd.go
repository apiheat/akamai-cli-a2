package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
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
	urlStr := fmt.Sprintf("%s/%s", URL, setID(c))

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := reportRespParse(data)
	errorCheck(err)

	printJSON(result)

	return nil
}

func resetProperty(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/%s/reset", URL, setID(c))

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

func errorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printJSON(str interface{}) {
	jsonRes, _ := json.MarshalIndent(str, "", "  ")
	fmt.Printf("%+v\n", string(jsonRes))
}

func fetchData(urlPath, method string, body io.Reader) (string, int) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, body)
	errorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	errorCheck(err)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt), resp.StatusCode
}

func setID(c *cli.Context) string {
	var id string
	if c.NArg() == 0 {
		log.Fatal("Please provide ID for property")
	}

	id = c.Args().Get(0)
	verifyID(id)
	return id
}

func verifyID(id string) {
	if _, err := strconv.Atoi(id); err != nil {
		errStr := fmt.Sprintf("Property ID should be number, you provided: %q\n", id)
		log.Fatal(errStr)
	}
}
