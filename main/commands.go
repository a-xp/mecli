package main

import (
	"github.com/urfave/cli"
	"MecomoApiMock/main/mecomo"
	"fmt"
)

func FetchDevices(ctx *cli.Context) error {
	fmt.Printf("Fetching all devices to local storage\n")
	api := mecomo.CreateClient(MECOMO_API_URL, ctx.String("api-user"), ctx.String("api-key"))
	fmt.Println(api.GetDevices(ctx.Int("limit")))
	return nil
}

func FetchTelemetry(ctx *cli.Context) error {
	return nil
}