package main

import (
	`fmt`
	`os`
	
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	
	`github.com/chaodoing/providers/containers`
)

func main() {
	err := os.Setenv("DIR", os.ExpandEnv("${PWD}"))
	if err != nil {
		panic(err)
	}
	boot, err := containers.Bootstrap("${DIR}/config/app.xml", true)
	if err != nil {
		panic(err)
	}
	
	err = boot.Handle(func(app *iris.Application, container *containers.Container) {
		app.Get(`/`, hero.Handler(func(ctx iris.Context, c containers.Container) {
			db, err := c.Db()
			if err != nil {
				c.Error(ctx, 3306, "数据库连接失败")
				return
			}
			var data map[string]interface{}
			err = db.Table("users").Find(&data).Error
			if err != nil {
				c.Error(ctx, 204, "数据查询失败")
				return
			}
			c.Success(ctx, data)
		}))
	}).Run()
	if err != nil {
		fmt.Println(err)
	}
}
