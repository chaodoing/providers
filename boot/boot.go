package boot

import (
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`github.com/kataras/iris/v12/middleware/logger`
	`github.com/kataras/iris/v12/middleware/recover`
	`os`
)

func Boot(conf string) Bootstrap {
	var config, err = XML(conf)
	if err != nil {
		panic(err)
	}
	var container = Container{
		conf: config,
	}
	container.db, err = container.MySQL()
	if err != nil {
		panic(err)
	}
	container.redis, err = container.Redis()
	hero.Register(container)
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Query:   true,
		Columns: true,
	}))
	drive, err := container.Log("iris-%Y-%m-%d.log")
	if err != nil {
		panic(err)
	}
	app.Logger().SetOutput(drive)
	if config.Log.Stdout {
		app.Logger().AddOutput(os.Stdout)
	}
	return Bootstrap{
		app:       app,
		container: container,
	}
}
