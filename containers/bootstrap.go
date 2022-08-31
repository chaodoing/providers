package containers

import (
	`fmt`
	`log`
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

type bootstrap struct {
	container *Container
	error     *lumberjack.Logger
	iris      *lumberjack.Logger
	app       *iris.Application
}

func (b *bootstrap) Handle(handle Handle) *bootstrap {
	b.app = iris.Default()
	handle(b.app, b.container)
	return b
}

func (b *bootstrap) Run() (err error) {
	b.error, err = b.container.logDriver(b.container.Log.Iris.Error, b.container.Log.Directory)
	if err != nil {
		return err
	}
	b.iris, err = b.container.logDriver(b.container.Log.Iris.File, b.container.Log.Directory)
	if err != nil {
		return
	}
	if b.container.Log.Console && b.container.Log.Record {
		b.app.Logger().SetOutput(b.iris)
		b.app.Logger().AddOutput(os.Stdout)
	} else if b.container.Log.Console {
		b.app.Logger().SetOutput(os.Stdout)
	} else {
		b.app.Logger().SetOutput(b.iris)
	}
	b.app.Logger().SetLevel(b.container.Log.Iris.Level)
	if b.container.Listener.Cors {
		b.app.AllowMethods(iris.MethodOptions)
		b.app.UseGlobal(middleware.Cors)
	}
	if fsutil.FileExist(b.container.Resource.Favicon) {
		b.app.Favicon(b.container.Resource.Favicon)
	}
	// 模板加载
	if fsutil.PathExist(b.container.Resource.Template.Directory) {
		b.app.RegisterView(iris.HTML(b.container.Resource.Template.Directory, b.container.Resource.Template.Extension))
	}
	// 静态目录
	if fsutil.PathExist(b.container.Resource.Asset.Directory) {
		b.app.HandleDir(b.container.Resource.Asset.Url, b.container.Resource.Asset.Directory)
	}
	
	// 上传目录
	if fsutil.PathExist(b.container.Resource.Upload.Directory) {
		b.app.HandleDir(b.container.Resource.Asset.Url, b.container.Resource.Asset.Directory)
	} else {
		if err := fsutil.Mkdir(b.container.Resource.Upload.Directory, 0755); err != nil {
			panic(err)
		} else {
			b.app.HandleDir(b.container.Resource.Upload.Url, b.container.Resource.Upload.Directory)
		}
	}
	b.app.Configure(iris.WithConfiguration(iris.Configuration{
		PostMaxMemory:       int64(b.container.Resource.Upload.Maximum) << 20,
		TimeFormat:          "2006-01-02 15:04:05",
		EnableOptimizations: true,
		Charset:             "UTF-8",
		DisableStartupLog:   !strings.HasPrefix(os.Getenv("ENV"), "dev"),
		Other: map[string]interface{}{
			"routes": b.app.GetRoutes(),
		},
	}))
	return b.app.Run(iris.Server(&http.Server{
		Addr:              fmt.Sprintf("%s:%d", b.container.Listener.Host, b.container.Listener.Port),
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Second * 6,
		WriteTimeout:      time.Second * 30,
		ErrorLog:          log.New(b.error, "", log.LstdFlags|log.Llongfile),
	}))
}

func Bootstrap(env string, isXml bool) (boot *bootstrap, err error) {
	var config *Env
	if isXml {
		config, err = XML(env)
		if err != nil {
			return
		}
	} else {
		config, err = JSON(env)
		if err != nil {
			return
		}
	}
	var container = Container{
		Env: config,
	}
	container.redis, err = container.Cache()
	if err != nil {
		return
	}
	container.db, err = container.Db()
	if err != nil {
		return
	}
	hero.Register(container)
	return &bootstrap{container: &container}, nil
}
