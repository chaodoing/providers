package main

import (
	`fmt`
	`os`
	
	`github.com/gookit/goutil/envutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	
	`github.com/chaodoing/providers/containers`
	`github.com/chaodoing/providers/http/controller`
	`github.com/chaodoing/providers/http/middleware`
)

const (
	// ENV     = "release"
	ENV     = "development"
	APP     = "providers"
	VERSION = "v2.0.0"
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
	boot, err := containers.Bootstrap("${DIR}/config/app.xml", true)
	if err != nil {
		panic(err)
	}
	
	err = boot.Handle(func(app *iris.Application, container *containers.Container) {
		app.Get(`/`, hero.Handler(middleware.Auth), hero.Handler(controller.Index))
		app.Post(`/login`, hero.Handler(controller.Login))
	}).Run()
	if err != nil {
		fmt.Println(err)
	}
}
