package containers

import (
	`github.com/kataras/iris/v12`
)

// Handle 路由方法
type Handle func(app *iris.Application, container *Container)
