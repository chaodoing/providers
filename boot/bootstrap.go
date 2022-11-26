package boot

import (
	`fmt`
	`github.com/gookit/goutil/fsutil`
	`github.com/kataras/iris/v12`
	`net/http`
	`os`
	`strings`
	`time`
)

type Bootstrap struct {
	config    *iris.Configuration
	server    *http.Server
	app       *iris.Application
	container Container
}

func (b Bootstrap) Configurator(config string) Bootstrap {
	irisConf := iris.YAML(os.ExpandEnv(config))
	b.config = &irisConf
	return b
}

func (b Bootstrap) Server(server *http.Server) Bootstrap {
	b.server = server
	return b
}

func (b Bootstrap) Handle(handle Handle) Bootstrap {
	handle(b.app)
	return b
}

func (b Bootstrap) Level(level string) Bootstrap {
	b.app.Logger().SetLevel(level)
	return b
}

func (b Bootstrap) Run() error {
	if b.container.conf.Station.AllowCrossDomain {
		b.app.AllowMethods(iris.MethodOptions)
		b.app.UseGlobal(Cors)
	}
	if fsutil.FileExist(b.container.conf.Static.Favicon) {
		b.app.Favicon(b.container.conf.Static.Favicon)
	}
	if fsutil.PathExist(b.container.conf.Template.Directory) {
		b.app.RegisterView(iris.HTML(b.container.conf.Template.Directory, b.container.conf.Template.Extension).Delims(b.container.conf.Template.Delimit.Left, b.container.conf.Template.Delimit.Right).Reload(b.container.conf.Template.Reload))
	}
	// 静态目录
	if fsutil.PathExist(b.container.conf.Static.Directory) {
		b.app.HandleDir(b.container.conf.Static.Url, b.container.conf.Static.Directory)
	}
	if fsutil.PathExist(b.container.conf.Upload.Directory) {
		b.app.HandleDir(b.container.conf.Upload.Url, b.container.conf.Upload.Directory)
	} else {
		if err := fsutil.Mkdir(b.container.conf.Upload.Directory, 0755); err != nil {
			panic(err)
		} else {
			b.app.HandleDir(b.container.conf.Upload.Url, b.container.conf.Upload.Directory)
		}
	}
	if b.server == nil {
		b.server = &http.Server{
			Addr:              fmt.Sprintf("%s:%d", b.container.conf.Station.Host, b.container.conf.Station.Port),
			ReadTimeout:       time.Second * 10,
			ReadHeaderTimeout: time.Second * 6,
			WriteTimeout:      time.Second * 30,
		}
	}
	if b.config == nil {
		b.config = &iris.Configuration{
			DisableStartupLog:   !strings.EqualFold(os.Getenv("ENV"), "development"),
			TimeFormat:          "2006-01-02 15:04:05",
			Charset:             "UTF-8",
			PostMaxMemory:       int64(b.container.conf.Upload.Maximum) << 20,
			EnableOptimizations: true,
			Other: map[string]interface{}{
				"routes": b.app.GetRoutes(),
			},
		}
	}
	return b.app.Run(iris.Server(b.server), iris.WithConfiguration(*b.config))
}
