package console

import (
	`encoding/xml`
)

type String struct {
	Comment  string `xml:"comment,attr"`
	XMLValue string `xml:",innerxml"`
}

type Delimit struct {
	Comment string `xml:"comment,attr"`
	Left    string `xml:"left,attr"`  // Left 左边变量分隔符号
	Right   string `xml:"right,attr"` // Right 右边变量分隔符号
}

type Station struct {
	Comment          string `xml:"comment,attr"`
	Host             String `xml:"host"`
	Port             String `xml:"port"`
	AllowCrossDomain String `xml:"allowCrossDomain"`
}

type Log struct {
	Comment   string `xml:"comment,attr"`
	Stdout    String `xml:"stdout"`
	Record    String `xml:"record"`
	Directory String `xml:"directory"`
}

type Database struct {
	Comment  string `xml:"comment,attr"`
	Host     String `xml:"host"`
	Port     String `xml:"port"`
	Db       String `xml:"db"`
	Username String `xml:"username"`
	Password String `xml:"password"`
	Charset  String `xml:"charset"`
	Record   String `xml:"record"`
	Level    String `xml:"level"`
}

type Redis struct {
	Comment string `xml:"comment,attr"`
	Host    String `xml:"host"`
	Port    String `xml:"port"`
	Db      String `xml:"db"`
	Auth    String `xml:"auth"`
	TTL     String `xml:"ttl"`
}

type Static struct {
	Comment   string `xml:"comment,attr"`
	Favicon   String `xml:"favicon"`
	Url       String `xml:"url"`
	Directory String `xml:"directory"`
}

type Template struct {
	Comment   string  `xml:"comment,attr"`
	Delimit   Delimit `xml:"delimit"`
	Directory String  `xml:"directory"`
	Extension String  `xml:"extension"`
	Reload    String  `xml:"reload"`
}

type Upload struct {
	Comment   string `xml:"comment,attr"`
	Maximum   String `xml:"maximum"`
	Url       String `xml:"url"`
	Directory String `xml:"directory"`
}

type configuration struct {
	XMLName xml.Name `xml:"app"`
	// Station 监听站点
	Station Station `xml:"station"`
	// Log 日志配置
	Log Log `xml:"log"`
	// Database 数据库配置
	Database Database `xml:"database"`
	// Redis 配置文件
	Redis Redis `xml:"redis"`
	// Static 静态资源文件配置
	Static Static `xml:"static"`
	
	// Template 模板目录配置
	Template Template `xml:"template"`
	
	// Upload 上传文件配置
	Upload Upload `xml:"upload"`
}

var Configuration = configuration{
	Station: Station{
		Comment:          "站点配置",
		Host:             String{Comment: "监听主机", XMLValue: "127.0.0.1"},
		Port:             String{Comment: "监听端口", XMLValue: "8080"},
		AllowCrossDomain: String{Comment: "允许跨域", XMLValue: "true"},
	},
	Log: Log{
		Comment:   "日志配置",
		Stdout:    String{Comment: "输出到控制台", XMLValue: "true"},
		Record:    String{Comment: "记录到日志文件", XMLValue: "true"},
		Directory: String{Comment: "日志存储目录", XMLValue: "${DIR}/logs"},
	},
	Database: Database{
		Comment:  "数据库配置",
		Host:     String{Comment: "连接主机", XMLValue: "127.0.0.1"},
		Port:     String{Comment: "连接端口", XMLValue: "3306"},
		Db:       String{Comment: "数据库名称", XMLValue: "test"},
		Username: String{Comment: "连接用户名", XMLValue: "test"},
		Password: String{Comment: "连接用户登录密码", XMLValue: ""},
		Charset:  String{Comment: "连接字符集", XMLValue: "utf8mb4"},
		Record:   String{Comment: "数据库日志文件名称", XMLValue: "mysql-%Y-%m-%d.log"},
		Level:    String{Comment: "数据库日志等级 1:silent 2:error 3:warn 4:info", XMLValue: "4"},
	},
	Redis: Redis{
		Comment: "Redis配置",
		Host:    String{Comment: "连接主机", XMLValue: "127.0.0.1"},
		Port:    String{Comment: "连接端口", XMLValue: "6379"},
		Db:      String{Comment: "数据库索引 default [0-15]", XMLValue: "0"},
		Auth:    String{Comment: "连接密码", XMLValue: ""},
		TTL:     String{Comment: "默认缓存时长[秒]", XMLValue: "6480000"},
	},
	Static: Static{
		Comment:   "静态目录配置文件",
		Favicon:   String{"网站图标", "${DIR}/resources/template/favicon.ico"},
		Url:       String{"访问地址", "/static"},
		Directory: String{"实际存储位置", "${DIR}/resource/template/static"},
	},
	Template: Template{
		Comment: "模板目录配置",
		Delimit: Delimit{
			Comment: "变量分隔符号",
			Left:    "{%",
			Right:   "%}",
		},
		Directory: String{"模板目录位置", "${DIR}/resources/template"},
		Extension: String{"扩展文件名称", ".html"},
		Reload:    String{"每次访问都加载配置", "false"},
	},
	Upload: Upload{
		Comment:   "上传配置",
		Maximum:   String{"允许上传的文件大小[MB]", "20"},
		Url:       String{"上传目录访问位置", "/upload"},
		Directory: String{"上传目录实际存储位置", "${DIR}/resources/upload"},
	},
}
