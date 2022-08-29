package containers

import (
	`errors`
	`fmt`
	`io`
	`log`
	`os`
	`path`
	`strings`
	`time`
	
	`github.com/go-redis/redis`
	`github.com/kataras/iris/v12`
	`github.com/lestrrat-go/strftime`
	`github.com/natefinch/lumberjack`
	`gorm.io/driver/mysql`
	`gorm.io/gorm`
	`gorm.io/gorm/logger`
	
	`github.com/chaodoing/providers/response`
)

type Container struct {
	*Env
	redis *redis.Client
	db    *gorm.DB
}

// logDriver 日志文件存储驱动
//  @param string name 日志文件名称
//  @param string directory 目录名称
func (c Container) logDriver(name, directory string) (data *lumberjack.Logger, err error) {
	var p *strftime.Strftime
	p, err = strftime.New(path.Join(directory, name))
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

// Db 初始化数据库
//  @return *gorm.DB db 数据库驱动
//  @return error err 错误
func (c Container) Db() (db *gorm.DB, err error) {
	if c.db != nil {
		return c.db, nil
	}
	var (
		logDrive    *lumberjack.Logger
		flag        = log.LstdFlags | log.Llongfile
		development = strings.HasPrefix(os.Getenv("ENV"), "dev")
		logs        logger.Interface
		write       io.Writer
	)
	logDrive, err = c.logDriver(c.Log.Database.File, c.Log.Directory)
	if err != nil {
		return
	}
	
	if c.Log.Console && c.Log.Record {
		write = io.MultiWriter(logDrive, os.Stdout)
	} else if c.Log.Console {
		write = os.Stdout
	} else {
		write = logDrive
	}
	if !development {
		flag = log.LstdFlags | log.Ldate | log.Ltime
	}
	// 开启日志
	if c.Log.Database.Enable {
		logs = logger.New(log.New(write, "", flag), logger.Config{
			Colorful: development,
			LogLevel: gormLevel(c.Log.Database.Level),
		})
	}
	drive, schema := c.dialect()
	if strings.EqualFold(drive, "mysql") {
		c.db, err = gorm.Open(mysql.Open(schema), &gorm.Config{
			NamingStrategy: nil,
			Logger:         logs,
			DryRun:         false,
			PrepareStmt:    true,
		})
		return c.db, err
	} else {
		return nil, errors.New("目前支持mysql")
	}
	return
}

// Cache 缓存数据
//  @return *redis.Client redisCli 缓存驱动
//  @return error err 错误
func (c Container) Cache() (redisCli *redis.Client, err error) {
	if c.redis != nil {
		return c.redis, nil
	}
	c.redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       c.Redis.Db,
	})
	var pong string
	pong, err = c.redis.Ping().Result()
	if err != nil || !strings.EqualFold(pong, "PONG") {
		return
	}
	return c.redis, nil
}

// Success 输出成功内容
func (c Container) Success(ctx iris.Context, data interface{}) {
	err := response.Responsive(ctx, data).Success()
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}

// Error 输出错误消息
func (c Container) Error(ctx iris.Context, status uint32, message string) {
	err := response.Responsive(ctx, nil).Error(status, message)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}

// Pagination 分页数据
func (c Container) Pagination(ctx iris.Context, pagination response.Pagination) {
	err := response.Responsive(ctx, pagination).Send()
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}
