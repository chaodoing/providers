package containers

import (
	`github.com/kataras/iris/v12`
)

type RouteDriver func(app *iris.Application)
