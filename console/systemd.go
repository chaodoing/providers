package console

import (
	`html/template`
	`os`
	
	`github.com/gookit/goutil/fsutil`
	`github.com/urfave/cli`
	
	`providers/assets`
)
var (
	dir string
	exec string
	config string
)
var Systemd = cli.Command{
	Name:        "systemd",
	ShortName:   "sys",
	Usage:       "生成Linux服务脚本",
	Description: "生成Linux [.service] 格式服务脚本",
	Category:    "框架命令",
	Flags:       []cli.Flag{
		cli.StringFlag{
			Name:        "dir,d",
			Usage:       "程序工作路径",
			Value:       os.Getenv("PWD"),
			Destination: &dir,
		},
		cli.StringFlag{
			Name:        "exec,e",
			Usage:       "可运行程序所在路径",
			Value:       os.Getenv("PWD") + "/bin/" + fsutil.Name(os.Args[0]),
			Destination: &exec,
		},
		cli.StringFlag{
			Name:        "config,c",
			Usage:       "配置文件问位置",
			Value:       os.ExpandEnv("${PWD}/config/app.xml"),
			Destination: &config,
		},
	},
	Action: func(c *cli.Context) error {
		var data = struct {
			App string
			Path string
			Execute string
			Config  string
		}{
			App: os.Getenv("APP"),
			Path: dir,
			Execute: exec,
			Config:  config,
		}
		service, err := assets.Asset("service/app.service")
		if err != nil {
			return err
		}
		tpl, err := template.New("systemd").Parse(string(service))
		if err != nil {
			return err
		}
		err = tpl.Execute(os.Stdout, data)
		if err != nil {
			return err
		}
		return nil
	},
}
