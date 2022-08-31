package main

import (
	`fmt`
	`os`
	
	`github.com/urfave/cli`
	
	`github.com/chaodoing/providers/console`
)

const (
	ENV     = "development"
	APP     = "providers"
	VERSION = "v2.0.0"
)

func main() {
	var err error
	err = os.Setenv("APP", APP)
	if err != nil {
		panic(err)
		return
	}
	err = os.Setenv("ENV", ENV)
	if err != nil {
		panic(err)
		return
	}
	err = os.Setenv("VERSION", VERSION)
	if err != nil {
		panic(err)
		return
	}
	
	var app = cli.NewApp()
	app.Name = APP
	app.Version = VERSION
	app.Description = "网站服务程序"
	app.Usage = "网站服务程序"
	app.Commands = []cli.Command{
		console.Systemd,
		console.Config,
		console.Env,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
