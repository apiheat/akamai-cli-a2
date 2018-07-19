# Akamai CLI for Adaptive Acceleration
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

### Local Install, if you choose not to use the akamai package manager
If you want to compile it from source, you will need Go 1.9 or later, and the [Glide](https://glide.sh) package manager installed:
1. Fetch the package:
   `go get https://github.com/partamonov/akamai-cli-a2`
1. Change to the package directory:
   `cd $GOPATH/src/github.com/partamonov/akamai-cli-a2`
1. Install dependencies using Glide:
   `glide install`
1. Compile the binary:
   `go build -ldflags="-s -w -X main.version=X.X.X" -o akamai-a2`

### Credentials
In order to use this configuration, you need to:
* Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Expects `default` section in .edgerc, can be changed via --section parameter

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

## Overview
The Adaptive Acceleration product uses Automatic Server Push, with the HTTP/2 protocol, and Automatic Preconnect to increase the speed of page loading. The Adaptive Acceleration API provides the ability to see which rules Adaptive Acceleration applies to a property. It also allows you to start the generation of new rules.

## Main Command Usage
```shell
NAME:
   akamai a2 - A CLI to interact with Akamai Adaptive Acceleration

USAGE:
   akamai a2 [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     report, r  Get a report for property [ID]
     reset, rm  Reset all existing info for property [ID]
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "${HOME}/.edgerc") [$AKAMAI_EDGERC]
   --debug                  Debug info
   --no-color               Disable color output
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

### Report command

You can run report for property

```shell
> akamai a2 report PROPERTY_ID

{
    "creationDate": "2016-08-25T16:45:02Z",
    "isActive": true,
    "lastModifiedDate": "2016-08-25T16:45:02Z",
    "lastReset": "2016-08-25T16:45:02Z",
    "zoneDeployDate": "2017-01-26T18:07:01Z",
    "version": 123,
    "commonPreconnectHeader": [
        "https://www.example0.com"
    ],
    "commonPushedResources": [
        "https://www.example.com/66dc044c-c886-4eb5-8ff2-d9721ec27fc3.js",
        "https://www.example.com/3e940cc8-8170-4e02-8ba4-5d967a997518.js",
        "https://www.example.com/1f682d27-2afa-4e3a-ac1a-2616629152e7.js",
        "https://www.example.com/12345678-c886-4eb5-8ff2-d9721ec27fc3.js"
    ],
    "pageSpecificRules": [
        {
            "basePageURL": "https://qa.www.example.com/checkout/receipt.jsp",
            "pageSpecificPreconnectHeader": [
                "https://www.example3.com"
            ],
            "pageSpecificPushedResources": [
                "https://qa.www.example.com/css/13455/all.css",
                "https://qa.www.example.com/css/13455/checkout/checkout.css",
                "https://qa.www.example.com/css/13455/checkout/receipt.css",
                "https://qa.www.example.com/js/13455/checkout/checkout.js",
                "https://qa.www.example.com/js/13455/be/ee522b75/scripts/lite-578e4f4264746d53-staging.js",
                "https://qa.www.example.com/js/13455/be/30ea2b75/scripts/lite-554041283961326d-staging.js",
                "https://qa.www.example.com/js/13455/be/b8eeeab7/scripts/lite-561eaa6565363300-staging.js"
            ],
            "basePageURL": "https://www.example.com/30dbc1a6-c9e4-474b-a525-70577c90c62d.html",
            "pageSpecificPreconnectHeader": [
                "https://www.example4.com"
            ],
            "pageSpecificPushedResources": [
                "https://www.example.com/c2be42fc-cc9a-4088-ab1c-1417ac212b6f.css",
                "https://www.example.com/c1ebcb4b-0a1a-4c64-9cc2-59de8ab25995.css",
                "https://www.example.com/34333a06-8e1e-4171-ae19-6560960d24df.css",
                "https://www.example.com/9eea3314-ea4e-4398-8dc7-a049c9706918.css",
                "https://www.example.com/8c8bf80a-dc13-4177-b360-3c7b97aa6d9b.css"
            ]
        }
    ]
}
```

#### Reset command


```shell
> akamai a2 reset PROPERTY_ID
...
```
