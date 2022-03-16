package containers

import (
	`errors`
	`fmt`
	`io`
	Logger `log`
	`os`
	`strings`
	`time`
	
	`github.com/go-redis/redis`
	`github.com/lestrrat-go/strftime`
	`github.com/natefinch/lumberjack`
	`gorm.io/driver/mysql`
	`gorm.io/gorm`
	`gorm.io/gorm/logger`
)

type Container struct {
	Config Config
	db     *gorm.DB
	rdx    *redis.Client
}

func (c *Container) Redis() (rdx *redis.Client, err error) {
	if c.rdx != nil {
		return c.rdx, nil
	}
	c.rdx = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", c.Config.Redis.Host, c.Config.Redis.Port),
		Password: c.Config.Redis.Auth,
		DB:       int(c.Config.Redis.DbIndex),
	})
	pong, err := c.rdx.Ping().Result()
	if err != nil || !strings.EqualFold(pong, "PONG") {
		return nil, errors.New("Error: redis连接失败")
	}
	return c.rdx, nil
}

// Logger 日志存储
func (c *Container) Logger(filename string) (data *lumberjack.Logger, err error) {
	p, err := strftime.New(os.ExpandEnv(filename))
	if err != nil {
		return
	}
	data = &lumberjack.Logger{
		Filename:   p.FormatString(time.Now()),
		MaxSize:    5,
		MaxAge:     1,
		MaxBackups: 31,
		LocalTime:  true,
		Compress:   true,
	}
	
	return
}

// Mysql 打开mysql连接
func (c *Container) Mysql() (db *gorm.DB, err error) {
	if c.db != nil {
		return c.db, nil
	}
	_, schema := c.Config.dialect()
	
	var (
		log *lumberjack.Logger
		Colorful bool
		logx logger.Interface
		out io.Writer
	)
	
	log, err = c.Logger(c.Get().Log.DbFile)
	if err != nil {
		return
	}
	
	if c.Get().Log.Console && c.Get().Log.OnSave {
		out = io.MultiWriter(os.Stdout, log)
	} else if c.Get().Log.Console {
		out = os.Stdout
		Colorful = true
	} else {
		out = log
	}
	logx = logger.New(Logger.New(out, "\r\n", Logger.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  gormLevel(c.Get().Log.DbLevel),
		IgnoreRecordNotFoundError: false,
		Colorful:                  Colorful,
	})
	c.db, err = gorm.Open(mysql.Open(schema), &gorm.Config{
		DryRun:      false,
		PrepareStmt: true,
		Logger:      logx,
	})
	if err != nil {
		return nil, err
	}
	return c.db, nil
}

// Authorized 用户认证
func (c *Container) Authorized() *Authorized {
	return &Authorized{
		rdx: c.rdx,
	}
}

// Get 获取配置文件
func (c *Container) Get() Config {
	return c.Config
}

// New 实例化容器
func New(config Config) Container {
	var (
		err       error
		container = Container{Config: config}
	)
	container.db, err = container.Mysql()
	if err != nil {
		panic(err)
	}
	container.rdx, err = container.Redis()
	if err != nil {
		panic(err)
	}
	return container
}