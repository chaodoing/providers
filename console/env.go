package console

import (
	`fmt`
	
	`github.com/urfave/cli`
	
	`github.com/chaodoing/providers/asset`
)

var Env = cli.Command{
	Name:        "env",
	ShortName:   "e",
	Usage:       "环境变量",
	Description: "生成环境变量文件",
	Category:    "Frame",
	Action: func(c *cli.Context) (err error) {
		content, err := asset.Asset("env/env")
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		return nil
	},
}
