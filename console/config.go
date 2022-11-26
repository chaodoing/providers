package console

import (
	`github.com/chaodoing/providers/o`
	`github.com/urfave/cli`
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
		return o.SaveXML(Configuration, file)
	},
}
