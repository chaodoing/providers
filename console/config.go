package console

import (
	`github.com/urfave/cli`
	
	`github.com/chaodoing/providers/console/environment`
)

var file string

var Config = cli.Command{
	Name:        "config",
	ShortName:   "c",
	Usage:       "生成Linux服务脚本",
	Description: "生成Linux [.service] 格式服务脚本",
	Category:    "Frame",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "file,f",
			Usage:       "配置文件输出位置",
			Required:    true,
			Value:       "./config.xml",
			Destination: &file,
		},
	},
	Action: func(c *cli.Context) (err error) {
		data, err := environment.InitEnvironment()
		if err != nil {
			panic(err)
		}
		return putil.SaveXML(data, file)
	},
}
