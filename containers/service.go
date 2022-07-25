package containers

import (
	"fmt"
	Logger "log"
	"net/http"
	"os"
	"strings"
	"time"
	
	"github.com/gookit/goutil/fsutil"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/natefinch/lumberjack"
	
	`github.com/chaodoing/providers/middleware`
)

func Service(container *Container, Driver RouteDriver) {
	var flag = Logger.LstdFlags
	if !strings.EqualFold(os.Getenv("ENVIRONMENT"), "development") || strings.EqualFold(os.Getenv("ENVIRONMENT"), "") {
		flag = Logger.Ldate | Logger.Ltime
	}
	Logger.SetFlags(flag)
	var (
		app               = iris.Default()
		err               error
		DisableStartupLog bool
		logger            *lumberjack.Logger
	)
	logger, err = container.Logger(container.Get().Log.Iris.File)
	if err != nil {
		panic(err)
		return
	}
	hero.Register(container)
	if container.Get().Log.Console && container.Get().Log.Console {
		app.Logger().SetOutput(logger)
		app.Logger().AddOutput(os.Stdout)
	}
	
	if !container.Get().Log.Console && container.Get().Log.IsSave {
		app.Logger().SetOutput(logger)
	}
	if container.Get().Log.Console && !container.Get().Log.IsSave {
		app.Logger().SetOutput(os.Stdout)
	}
	app.Logger().SetLevel(irisLevel(container.Get().Log.Iris.Level))
	
	if strings.EqualFold(os.Getenv("ENVIRONMENT"), "release") {
		DisableStartupLog = true
	}
	if container.Get().App.Cross {
		app.AllowMethods(iris.MethodOptions)
		app.UseGlobal(middleware.CrossAllowed)
	}
	
	// Favicon 图标
	if fsutil.FileExist(os.ExpandEnv(container.Config.AssetBundle.Favicon)) {
		app.Favicon(os.ExpandEnv(container.Config.AssetBundle.Favicon))
	}
	// 模板加载
	if fsutil.PathExist(os.ExpandEnv(container.Config.AssetBundle.Template)) {
		app.RegisterView(iris.HTML(os.ExpandEnv(container.Config.AssetBundle.Template), ".html"))
	}
	// 静态目录
	if fsutil.PathExist(os.ExpandEnv(container.Config.AssetBundle.Static.Path)) {
		app.HandleDir(container.Config.AssetBundle.Static.Uri, os.ExpandEnv(container.Config.AssetBundle.Static.Path))
	}
	// 上传目录
	if fsutil.PathExist(os.ExpandEnv(container.Config.AssetBundle.Upload.Path)) {
		app.HandleDir(container.Config.AssetBundle.Upload.Uri, os.ExpandEnv(container.Config.AssetBundle.Upload.Path))
	} else {
		if err := fsutil.Mkdir(os.ExpandEnv(container.Config.AssetBundle.Upload.Path), 0755); err != nil {
			panic(err)
		} else {
			app.HandleDir(container.Config.AssetBundle.Upload.Uri, os.ExpandEnv(container.Config.AssetBundle.Upload.Path))
		}
	}
	Driver(app, container)
	app.Configure(iris.WithConfiguration(iris.Configuration{
		PostMaxMemory:       int64(container.Config.AssetBundle.Upload.Maximum) << 20,
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
