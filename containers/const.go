package containers

import (
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
)

type RouteDriver func(app *iris.Application, container *Container, dim *hero.Hero)
