package environment

import (
	`encoding/xml`
	`fmt`
	
	`github.com/chaodoing/providers/console/environment/database`
	`github.com/chaodoing/providers/console/environment/listener`
	`github.com/chaodoing/providers/console/environment/log`
	`github.com/chaodoing/providers/console/environment/redis`
	`github.com/chaodoing/providers/console/environment/resource`
)

type Environment struct {
	XMLName  xml.Name           `json:"-" xml:"root"`
	Listener *listener.Listener `xml:"listener"`
	Log      *log.Log           `xml:"log"`
	Database *database.Database `xml:"database"`
	Redis    *redis.Redis       `xml:"redis"`
	Resource *resource.Resource `xml:"resource"`
}

func InitEnvironment() (data *Environment, err error) {
	var (
		Listener *listener.Listener
		Log      *log.Log
		Database *database.Database
		Redis    *redis.Redis
		Resource *resource.Resource
	)
	fmt.Println("网站配置")
	Listener, err = listener.NewListener()
	if err != nil {
		return
	}
	fmt.Println("日志配置")
	Log, err = log.NewLog()
	if err != nil {
		return
	}
	fmt.Println("数据库配置")
	Database, err = database.NewDatabase()
	if err != nil {
		return
	}
	fmt.Println("Redis配置")
	Redis, err = redis.NewRedis()
	if err != nil {
		return
	}
	fmt.Println("资源配置")
	Resource, err = resource.NewResource()
	if err != nil {
		return
	}
	data = &Environment{
		Listener: Listener,
		Log:      Log,
		Database: Database,
		Redis:    Redis,
		Resource: Resource,
	}
	return
}
