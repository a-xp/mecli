package main

import (
	"github.com/urfave/cli"
	"MecomoApiMock/main/mecomo"
	"fmt"
	dao2 "MecomoApiMock/main/dao"
)

func FetchDevices(ctx *cli.Context) error {
	fmt.Printf("Fetching all devices to local storage\n")
	api := mecomo.CreateClient(MECOMO_API_URL, ctx.String("api-user"), ctx.String("api-key"))
	devices :=api.GetDevices(ctx.Int("limit"))
	dao,err:= dao2.CreateDAO(ctx.GlobalString("dsn"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dao.Close()
	dao.StoreDevices(devices)
	fmt.Printf("%d devices were stored in DB", len(devices))
	return nil
}

func FetchTelemetry(ctx *cli.Context) error {
	return nil
}

func PushDevices(ctx *cli.Context) error {
	return nil
}