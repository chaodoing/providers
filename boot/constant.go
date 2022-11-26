package boot

import (
	`github.com/kataras/iris/v12`
)

type Handle func(app *iris.Application)

const (
	Basic  = "Basic "
	Bearer = "Bearer "
)
