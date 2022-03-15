package containers

import (
	`fmt`
	`net/http`
	`os`
	`strings`
	`time`
	
	`github.com/gookit/goutil/fsutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`github.com/natefinch/lumberjack`
	
	`github.com/chaodoing/providers/middleware`
)

func Service(container *Container, Driver RouteDriver) {
	var (
		app               = iris.Default()
		err               error
		DisableStartupLog bool
		logger            *lumberjack.Logger
	)
	logger, err = container.Logger(container.Get().Log.IrisFile)
	if err != nil {
		panic(err)
		return
	}
	hero.Register(container)
	if container.Get().Log.Console && container.Get().Log.OnSave {
		app.Logger().SetOutput(logger)
		app.Logger().AddOutput(os.Stdout)
	}
	if !container.Get().Log.Console && container.Get().Log.OnSave {
		app.Logger().SetOutput(logger)
	}
	if container.Get().Log.Console && !container.Get().Log.OnSave {
		app.Logger().SetOutput(os.Stdout)
	}
	app.Logger().SetLevel(irisLevel(container.Get().Log.Level))
	
	if strings.EqualFold(os.Getenv("ENVIRONMENT"), "release") {
		DisableStartupLog = true
	}
	if container.Config.App.AllowedCross {
		app.AllowMethods(iris.MethodOptions)
		app.UseGlobal(middleware.AllowedCrossDomain)
	}
	
	// Favicon 图标
	if fsutil.FileExist(os.ExpandEnv(container.Config.AssetBundle.Favicon)) {
		app.Favicon(os.ExpandEnv(container.Config.AssetBundle.Favicon))
	}
	// 模板加载
	if fsutil.PathExist(container.Config.AssetBundle.Template) {
		app.RegisterView(iris.HTML(os.ExpandEnv(container.Config.AssetBundle.Template), ".html"))
	}
	// 静态目录
	if fsutil.PathExist(container.Config.AssetBundle.Static) {
		app.HandleDir(container.Config.AssetBundle.StaticUri, os.ExpandEnv(container.Config.AssetBundle.Static))
	}
	// 上传目录
	if fsutil.PathExist(container.Config.AssetBundle.Upload) {
		app.HandleDir(container.Config.AssetBundle.UploadUri, os.ExpandEnv(container.Config.AssetBundle.Upload))
	} else {
		if err := fsutil.Mkdir(os.ExpandEnv(container.Config.AssetBundle.Upload), 0755); err != nil {
			panic(err)
		} else {
			app.HandleDir(container.Config.AssetBundle.UploadUri, os.ExpandEnv(container.Config.AssetBundle.Upload))
		}
	}
	Driver(app)
	app.Configure(iris.WithConfiguration(iris.Configuration{
		PostMaxMemory:       int64(container.Config.AssetBundle.UploadMaximum) << 20,
		TimeFormat:          "2006-01-02 15:04:05",
		EnableOptimizations: true,
		Charset:             "UTF-8",
		DisableStartupLog:   DisableStartupLog,
		Other: map[string]interface{}{
			"routes": app.GetRoutes(),
		},
	}))
	err = app.Run(iris.Server(&http.Server{
		Addr:              fmt.Sprintf("%s:%v", container.Config.App.Host, container.Config.App.Port),
		ReadTimeout:       time.Second * 10,
		WriteTimeout:      time.Second * 30,
		ReadHeaderTimeout: time.Second * 6,
	}), iris.WithoutInterruptHandler)
	if err != nil {
		panic(err)
	}
}
