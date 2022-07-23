package containers

import (
	"errors"
	"fmt"
	"io"
	Logger "log"
	"os"
	"strings"
	"time"
	
	"github.com/go-redis/redis"
	"github.com/lestrrat-go/strftime"
	"github.com/natefinch/lumberjack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gromlogger "gorm.io/gorm/logger"
)

type Provides map[string]interface{}

type Container struct {
	Config   Config
	db       *gorm.DB
	rdx      *redis.Client
	provides Provides
}

func (c *Container) Redis() (rdx *redis.Client, err error) {
	if c.rdx != nil {
		return c.rdx, nil
	}
	c.rdx = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", c.Config.Redis.Host, c.Config.Redis.Port),
		Password: c.Config.Redis.Auth,
		DB:       int(c.Config.Redis.Index),
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
		log      *lumberjack.Logger
		Colorful bool
		logx     gromlogger.Interface
		out      io.Writer
	)
	
	log, err = c.Logger(c.Get().Log.Database.File)
	if err != nil {
		return
	}
	
	if c.Get().Log.Console && c.Get().Log.Console {
		out = io.MultiWriter(os.Stdout, log)
	} else if c.Get().Log.Console {
		out = os.Stdout
		Colorful = true
	} else {
		out = log
	}
	if strings.EqualFold(os.Getenv("ENVIRONMENT"), "") {
		err := os.Setenv("ENVIRONMENT", "development")
		if err != nil {
			panic(err)
		}
	}
	var flag = Logger.LstdFlags | Logger.Llongfile
	if !strings.EqualFold(os.Getenv("ENVIRONMENT"), "development") {
		flag = Logger.Ldate | Logger.Ltime
	}
	logx = NewGormLog(Logger.New(out, "[GORM] ", flag), gromlogger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  gormLevel(c.Get().Log.Database.Level),
		IgnoreRecordNotFoundError: false,
		Colorful:                  Colorful,
	})
	if !c.Get().Log.Database.Enable {
		logx = nil
	}
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
		rdx:    c.rdx,
		Expire: int64(c.Config.Redis.Expire),
	}
}

// Provides 获取服务内容
func (c *Container) Provides() Provides {
	return c.provides
}

func (c *Container) AddProvide(key string, value interface{}) {
	c.provides[key] = value
}

func (c *Container) GetProvide(key string) (interface{}, error) {
	if data, ok := c.provides[key]; ok {
		return data, nil
	}
	return nil, errors.New("对象找不到")
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
	container.provides = make(Provides)
	container.db, err = container.Mysql()
	if err != nil {
		panic(err)
	}
	if !strings.EqualFold(container.Config.Redis.Host, "") {
		container.rdx, err = container.Redis()
		if err != nil {
			panic(err)
		}
	}
	return container
}
