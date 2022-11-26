package boot

import (
	`encoding/json`
	`encoding/xml`
	`fmt`
	`github.com/go-ini/ini`
	`os`
	`strings`
)

type Configuration struct {
	XMLName xml.Name `json:"-" xml:"app"`
	Station struct {
		Host             string `json:"host" xml:"host" ini:"Host" field:"监听主机"`                                     // Host 监听主机
		Port             uint16 `json:"port" xml:"port" ini:"Port" field:"监听端口"`                                     // Port 监听端口
		AllowCrossDomain bool   `json:"allowCrossDomain" xml:"allowCrossDomain" ini:"AllowCrossDomain" field:"允许跨域"` // AllowCrossDomain 允许跨域
	} `json:"station" xml:"station" ini:"Station" field:"监听站点"` // Station 监听站点
	Log struct {
		Stdout    bool   `json:"stdout" xml:"stdout" ini:"Stdout" field:"输出到控制台"`          // Stdout 输出到控制台
		Record    bool   `json:"record" xml:"record" ini:"Record" field:"记录到日志文件"`        // Record 记录到日志文件
		Directory string `json:"directory" xml:"directory" ini:"Directory" field:"日志存储目录"` // Directory 日志存储目录
	} `json:"log" xml:"log" ini:"Log" field:"日志配置"` // Log 日志配置
	Database struct {
		Host     string `json:"host" xml:"host" ini:"Host" field:"连接主机"`               // Host 连接主机
		Port     uint16 `json:"port" xml:"port" ini:"Port" field:"连接端口"`               // Port 连接端口
		Db       string `json:"db" xml:"db" ini:"Db" field:"数据库名称"`                   // Db 数据库名称
		Username string `json:"username" xml:"username" ini:"Username" field:"连接用户名"` // Username 连接用户名
		Password string `json:"password" xml:"password" ini:"Password" field:"连接密码"`   // Password 连接密码
		Charset  string `json:"charset" xml:"charset" ini:"Charset" field:"连接字符集"`    // Charset 连接字符集
		Record   string `json:"record" xml:"record" ini:"Record" field:"记录到日志文件"`   // Record 记录到日志文件
		Level    int    `json:"level" xml:"level" ini:"Level"  field:"日志等级"`           // Level 日志等级
	} `json:"database" xml:"database" field:"数据库配置"` // Database 数据库配置
	Redis struct {
		Host string `json:"host" xml:"host" ini:"Host"` // Host 连接主机
		Port uint16 `json:"port" xml:"port" ini:"Port"` // Port 连接端口
		Db   int    `json:"db" xml:"db" ini:"Db"`       // Db 数据库位置
		Auth string `json:"auth" xml:"auth" ini:"Auth"` // Auth 连接密码
		TTL  uint64 `json:"ttl" xml:"ttl" ini:"TTL"`    // TTL 缓存时长
	} `json:"redis" xml:"redis" ini:"Redis"` // Redis 配置文件
	
	// Static 静态资源文件配置
	Static struct {
		Favicon   string `json:"favicon" xml:"favicon" ini:"Favicon"`       // Favicon 网站图标配置
		Url       string `json:"url" xml:"url" ini:"Url"`                   // Url 访问路径
		Directory string `json:"directory" xml:"directory" ini:"Directory"` // Directory 静态资源位置
	} `json:"static" xml:"static" ini:"Static"`
	
	// Template 模板目录配置
	Template struct {
		Delimit struct {
			Left  string `json:"left" xml:"left,attr" ini:"Left"`    // Left 左边变量分隔符号
			Right string `json:"right" xml:"right,attr" ini:"Right"` // Right 右边变量分隔符号
		} `json:"delimit" xml:"delimit" ini:"Template.Delimit"`             // Delimit 变量边界符号
		Directory string `json:"directory" xml:"directory" ini:"Directory"` // Directory 模板目录位置
		Extension string `json:"extension" xml:"extension" ini:"Extension"` // Extension 模板文件扩展名称
		Reload    bool   `json:"reload" xml:"reload" ini:"Reload"`          // Reload 每次都重新加载模板
	} `json:"template" xml:"template" ini:"Template"`
	
	// Upload 上传文件配置
	Upload struct {
		Maximum   uint64 `json:"maximum" xml:"maximum" ini:"Maximum"`       // Maximum 允许上传的文件最大大小 单位 MB
		Url       string `json:"url" xml:"url" ini:"Url"`                   // Url 上传访问路径
		Directory string `json:"directory" xml:"directory" ini:"Directory"` // Directory 上传目录实际位置
	} `json:"upload" xml:"upload" ini:"Upload"`
}

// LoadEnv 加载环境变量
func (c *Configuration) LoadEnv() *Configuration {
	config, err := ini.Load(os.ExpandEnv("${DIR}/.env"))
	if err != nil {
		return nil
	}
	err = config.MapTo(c)
	if err != nil {
		return nil
	}
	return c
}

// Dialect 转换数据配置
//	@return schema=root:123.com@tcp(127.0.0.1:3306)/arrangement?charset=utf8mb4&parseTime=True&loc=Local
func (c *Configuration) Dialect() (schema string) {
	schema = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Db, c.Database.Charset)
	return
}

// parseDirectory 解析目录
func (c *Configuration) parseDirectory() *Configuration {
	c.Log.Directory = strings.TrimRight(os.ExpandEnv(c.Log.Directory), string(os.PathSeparator))
	c.Static.Favicon = os.ExpandEnv(c.Static.Favicon)
	c.Template.Directory = strings.TrimRight(os.ExpandEnv(c.Template.Directory), string(os.PathSeparator))
	c.Static.Directory = strings.TrimRight(os.ExpandEnv(c.Static.Directory), string(os.PathSeparator))
	c.Upload.Directory = strings.TrimRight(os.ExpandEnv(c.Upload.Directory), string(os.PathSeparator))
	return c
}

// XML 加载配置文件
func XML(env string) (data *Configuration, err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(env))
	if err != nil {
		return
	}
	err = xml.Unmarshal(content, &data)
	if err != nil {
		return
	}
	c := data.LoadEnv()
	if c != nil {
		data = c
	}
	data = data.parseDirectory()
	return
}

// JSON 加载配置文件
func JSON(env string) (data *Configuration, err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(env))
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return
	}
	c := data.LoadEnv()
	if c != nil {
		data = c
	}
	data = data.parseDirectory()
	return
}
