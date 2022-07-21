package main

import (
	"os"

	"github.com/urfave/cli"

	"virtualization/providers/console"
)

var (
	ENVIRONMENT = "development"
	VERSION     = "v1.0.0"
	NAME        = "wxamp"
)

func main() {
	os.Setenv("ENVIRONMENT", ENVIRONMENT)
	os.Setenv("VERSION", VERSION)
	os.Setenv("APP", NAME)

	app := cli.NewApp()
	app.Name = NAME
	app.Version = VERSION
	app.Commands = []cli.Command{
		console.Systemd,
		console.Config,
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
