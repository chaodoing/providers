# 修改框架基础结构

```go
package main

import (
	`github.com/chaodoing/providers/boot`
	`github.com/chaodoing/providers/models`
	`github.com/chaodoing/providers/o`
	`github.com/gookit/goutil/envutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`log`
	`os`
)

const (
	// ENV     = "production"
	ENV     = "development"
	APP     = "providers"
	VERSION = "v1.0.0"
)

func main() {
	var err error
	err = os.Setenv("APP", envutil.Getenv("APP", APP))
	if err != nil {
		panic(err)
		return
	}
	err = os.Setenv("ENV", envutil.Getenv("ENV", ENV))
	if err != nil {
		panic(err)
		return
	}
	err = os.Setenv("VERSION", envutil.Getenv("VERSION", VERSION))
	if err != nil {
		panic(err)
		return
	}
	
	err = os.Setenv("DIR", os.ExpandEnv("${PWD}"))
	if err != nil {
		panic(err)
	}
	err = boot.Boot("${DIR}/config/app.xml").Handle(func(app *iris.Application) {
		app.Get(`/`, hero.Handler(func(ctx iris.Context, container boot.Container) {
			auth, err := container.Auth()
			if err != nil {
				o.O(ctx, 1, err.Error(), nil)
				return
			}
			db, err := container.MySQL()
			if err != nil {
				log.Fatal(err)
				return
			}
			data, err := models.CategoriesMgr(db.Order("`sort` ASC")).Gets()
			if err != nil {
				o.O(ctx, 1, err.Error(), nil)
				return
			}
			err = auth.Create(ctx, data)
			if err != nil {
				o.O(ctx, 1, err.Error(), nil)
				return
			}
			o.O(ctx, 0, "Ok", data)
		}))
	}).Run()
	if err != nil {
		log.Fatal(err)
	}
}
```