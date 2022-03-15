package console

import (
	`fmt`
	`os`
	
	`github.com/urfave/cli`
	
	`providers/assets`
)

var (
	isXML  bool = true
	isJSON bool = false
)
var Config = cli.Command{
	Name:        "config",
	ShortName:   "conf",
	Usage:       "默认配置文件生成",
	Description: "生成默认的配置文件",
	Category:    "框架命令",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "xml",
			Usage:       "生成xml配置文件",
			FilePath:    "",
			Required:    false,
			Destination: &isXML,
		},
		cli.BoolFlag{
			Name:        "json",
			Usage:       "生成json配置文件",
			FilePath:    "",
			Required:    false,
			Destination: &isJSON,
		},
	},
	Action: func(c *cli.Context) error {
		var (
			config []byte
			err    error
		)
		if isXML {
			config, err = assets.Asset("config/app.xml")
		} else if isJSON {
			config, err = assets.Asset("config/app.json")
		} else {
			config, err = assets.Asset("config/app.xml")
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(config))
		return nil
	},
}
