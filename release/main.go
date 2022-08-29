package main

import (
	`fmt`
	`os`
	
	`github.com/urfave/cli`
	
	`github.com/chaodoing/providers/console`
)

func main() {
	var app = cli.NewApp()
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
