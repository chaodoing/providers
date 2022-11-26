package console

import (
	`bytes`
	`github.com/chaodoing/providers/assets`
	`github.com/chaodoing/providers/console/systemd`
	`github.com/gookit/goutil/fsutil`
	`github.com/urfave/cli`
	`html/template`
)

var Systemd = cli.Command{
	Name:        "systemd",
	ShortName:   "s",
	Usage:       "生成Linux服务脚本",
	Description: "生成Linux [.service] 格式服务脚本",
	Category:    "Frame",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "file,f",
			Usage:       "服务脚本输出位置",
			Required:    true,
			Value:       "./app.service",
			Destination: &file,
		},
	},
	Action: func(c *cli.Context) (err error) {
		var sys systemd.Systemd
		sys, err = systemd.NewSystemd()
		if err != nil {
			return
		}
		content, err := assets.Asset("service/app.service")
		if err != nil {
			return
		}
		tpl, err := template.New("systemd").Parse(string(content))
		if err != nil {
			return
		}
		buf := new(bytes.Buffer)
		err = tpl.Execute(buf, sys)
		_, err = fsutil.PutContents(file, buf.String())
		return
	},
}
