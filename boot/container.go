package boot

import (
	`errors`
	`fmt`
	`github.com/go-redis/redis`
	`github.com/lestrrat-go/strftime`
	`github.com/natefinch/lumberjack`
	`gorm.io/driver/mysql`
	`gorm.io/gorm`
	`gorm.io/gorm/logger`
	`io`
	`log`
	`os`
	`path`
	`strings`
	`time`
)

type Container struct {
	conf  *Configuration
	redis *redis.Client
	db    *gorm.DB
}

func (c Container) MySQL() (db *gorm.DB, err error) {
	if c.db != nil {
		return c.db, nil
	}
	var (
		record logger.Interface
		drive  *lumberjack.Logger
		write  io.Writer
	)
	// 如果开启数据库日志记录
	if !strings.EqualFold(c.conf.Database.Record, "") {
		drive, err = c.Log(c.conf.Database.Record)
		if err != nil {
			return
		}
		// 输出到控制台同时输出到日志文件
		if c.conf.Log.Stdout {
			write = io.MultiWriter(drive, os.Stdout)
		} else { // 只输出到文件
			write = drive
		}
	} else {
		err = errors.New("数据库日志记录文件不能为空")
		return
	}
	
	flag := log.LstdFlags | log.Ldate | log.Ltime
	record = logger.New(log.New(write, "", flag), logger.Config{
		Colorful: false,
		LogLevel: logger.LogLevel(c.conf.Database.Level),
	})
	schema := c.conf.Dialect()
	c.db, err = gorm.Open(mysql.Open(schema), &gorm.Config{
		SkipDefaultTransaction: true,   // SkipDefaultTransaction 跳过默认事务
		FullSaveAssociations:   true,   // FullSaveAssociations 在创建或更新时，是否更新关联数据
		Logger:                 record, // Logger 日志接口，用于实现自定义日志
		DryRun:                 false,  // DryRun 生成 SQL 但不执行，可以用于准备或测试生成的 SQL
		PrepareStmt:            true,   // PrepareStmt 是否禁止创建 prepared statement 并将其缓存
		AllowGlobalUpdate:      false,  // AllowGlobalUpdate 是否允许全局 update/delete
		QueryFields:            true,   // QueryFields 执行查询时，是否带上所有字段
	})
	return c.db, err
}

func (c Container) Redis() (rds *redis.Client, err error) {
	if c.redis != nil {
		return c.redis, nil
	}
	c.redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.conf.Redis.Host, c.conf.Redis.Port),
		Password: c.conf.Redis.Auth,
		DB:       c.conf.Redis.Db,
	})
	var pong string
	pong, err = c.redis.Ping().Result()
	if err != nil || !strings.EqualFold(pong, "PONG") {
		return
	}
	return c.redis, nil
}

// Log 日志文件存储驱动
//	@param string name 日志文件名称
//	@param string directory 目录名称
func (c Container) Log(name string) (drive *lumberjack.Logger, err error) {
	var p *strftime.Strftime
	p, err = strftime.New(path.Join(c.conf.Log.Directory, name))
	drive = &lumberjack.Logger{
		Filename:   p.FormatString(time.Now()),
		MaxSize:    5,
		MaxAge:     1,
		MaxBackups: 31,
		LocalTime:  true,
		Compress:   true,
	}
	return
}

// Auth 用户认证
func (c Container) Auth() (auth Auth, err error) {
	var redisCli *redis.Client
	redisCli, err = c.Redis()
	if err != nil {
		return
	}
	return Auth{redisCli: redisCli, TTL: time.Duration(c.conf.Redis.TTL * uint64(time.Second))}, nil
}
