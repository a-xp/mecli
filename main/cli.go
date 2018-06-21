package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
)

const MECOMO_API_URL  = "http://www.mecfleet.com/API/"

func main() {
	app:= cli.NewApp()
	app.Name = "Mecomo API test utils"
	app.Description = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "access-key", Usage: "AWS access key"},
		cli.StringFlag{Name: "secret-key", Usage: "AWS secret key"},
	}
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name: "fetch-devices",
			Aliases: []string{"fd"},
			Usage: "Fetch devices from provider API",
			Action: FetchDevices,
			Flags: []cli.Flag{
				cli.IntFlag{Name: "limit", Usage: "Max number of devices to fetch", Value: 100},
				cli.StringFlag{Name: "api-key", Usage: "Provider API key"},
				cli.StringFlag{Name: "api-user", Usage: "Provider API user"},
			},
		},
	}

	app.Action = func (ctx *cli.Context) error {
		fmt.Printf("Executing %s\n", ctx.Command.Name)
		return nil
	}

	app.Run(os.Args)
}
