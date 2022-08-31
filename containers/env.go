package containers

import (
	`encoding/json`
	`encoding/xml`
	`fmt`
	`os`
	`strings`
	
	`github.com/go-ini/ini`
)

// Env 配置内容
type Env struct {
	XMLName xml.Name `json:"-" xml:"root"`
	// Listener 网站监听
	Listener struct {
		Host string `json:"host" xml:"host" ini:"Host"` // Host 监听主机
		Port uint16 `json:"port" xml:"port" ini:"Port"` // Port 监听端口
		Cors bool   `json:"cors" xml:"cors" ini:"Cors"` // Cors 允许跨域
	} `json:"listener" xml:"listener" ini:"Listener"`
	// Log 日志配置
	Log struct {
		Console   bool   `json:"console" xml:"console" ini:"Console"`       // Console 输出到控制台
		Record    bool   `json:"record" xml:"record" ini:"Record"`          // Record 记录到日志文件
		Directory string `json:"directory" xml:"directory" ini:"Directory"` // Directory 日志存储目录
		// Iris 框架日志配置
		Iris struct {
			Level string `json:"level" xml:"level" ini:"Level"` // Level 日志等级
			File  string `json:"file" xml:"file" ini:"File"`    // File 日志文件
			Error string `json:"error" xml:"error" ini:"Error"` // Error 错误日志文件
		} `json:"iris" xml:"iris" ini:"Log.Iris"`
		// Database 数据库日志
		Database struct {
			Enable bool   `json:"enable" xml:"enable" ini:"Enable"` // Enable 启用日志
			Level  string `json:"level" xml:"level" ini:"Level"`    // Level 日志等级
			File   string `json:"file" xml:"file" ini:"File"`       // File 数据库日志文件
		} `json:"database" xml:"database" ini:"Log.Database"`
	} `json:"log" xml:"log" ini:"Log"`
	
	// Database 连接数据库
	Database struct {
		Drive    string `json:"drive" xml:"drive" ini:"Drive"`          // Drive 数据库类型
		Host     string `json:"host" xml:"host" ini:"Host"`             // Host 连接主机
		Port     uint16 `json:"port" xml:"port" ini:"Port"`             // Port 连接端口
		Db       string `json:"db" xml:"db" ini:"Db"`                   // Db 数据库名称
		Username string `json:"username" xml:"username" ini:"Username"` // Username 连接用户名
		Password string `json:"password" xml:"password" ini:"Password"` // Password 连接密码
		Charset  string `json:"charset" xml:"charset" ini:"Charset"`    // Charset 连接字符集
	} `json:"database" xml:"database" ini:"Database"`
	// Cache redis配置
	Redis struct {
		Host     string `json:"host" xml:"host" ini:"Host"`             // Host 连接主机
		Port     uint16 `json:"port" xml:"port" ini:"Port"`             // Port 连接端口
		Db       int    `json:"db" xml:"db" ini:"Db"`                   // Db 数据库位置
		Password string `json:"password" xml:"password" ini:"Password"` // Password 连接密码
		TTL      uint64 `json:"ttl" xml:"ttl" ini:"TTL"`                // TTL 缓存时长
	} `json:"redis" xml:"redis" ini:"Redis"`
	
	// Resource 资源文件配置
	Resource struct {
		Favicon string `json:"favicon" xml:"favicon" ini:"Favicon"` // Favicon 网站图标配置
		// Template 模板目录配置
		Template struct {
			Directory string `json:"directory" xml:"directory" ini:"Directory"` // Directory 模板目录位置
			Extension string `json:"extension" xml:"extension" ini:"Extension"` // Extension 模板文件扩展名称
			Reload    bool   `json:"reload" xml:"reload" ini:"Reload"`          // Reload 每次都重新加载模板
		} `json:"template" xml:"template" ini:"Resource.Template"`
		// Asset 静态资源数据配置
		Asset struct {
			Url       string `json:"url" xml:"url" ini:"Url"`                   // Url 访问路径
			Directory string `json:"directory" xml:"directory" ini:"Directory"` // Directory 静态资源位置
		} `json:"asset" xml:"asset" ini:"Resource.Asset"`
		// Upload 上传文件配置
		Upload struct {
			Maximum   uint64 `json:"maximum" xml:"maximum" ini:"Maximum"`       // Maximum 允许上传的文件最大大小 单位 MB
			Url       string `json:"url" xml:"url" ini:"Url"`                   // Url 上传访问路径
			Directory string `json:"directory" xml:"directory" ini:"Directory"` // Directory 上传目录实际位置
		} `json:"upload" xml:"upload" ini:"Resource.Upload"`
	} `json:"resource" xml:"resource" ini:"Resource"`
}

// LoadEnv 加载环境变量
func (c *Env) LoadEnv() *Env {
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

// parseDirectory 解析目录
func (c *Env) parseDirectory() *Env {
	c.Log.Directory = strings.TrimRight(os.ExpandEnv(c.Log.Directory), string(os.PathSeparator))
	c.Resource.Favicon = os.ExpandEnv(c.Resource.Favicon)
	c.Resource.Template.Directory = strings.TrimRight(os.ExpandEnv(c.Resource.Template.Directory), string(os.PathSeparator))
	c.Resource.Asset.Directory = strings.TrimRight(os.ExpandEnv(c.Resource.Asset.Directory), string(os.PathSeparator))
	c.Resource.Upload.Directory = strings.TrimRight(os.ExpandEnv(c.Resource.Upload.Directory), string(os.PathSeparator))
	return c
}

// dialect 转换数据配置
//	@return dialect=mysql
//	@return schema=root:123.com@tcp(127.0.0.1:3306)/arrangement?charset=utf8mb4&parseTime=True&loc=Local
//	@return logMode=false
func (c *Env) dialect() (dialect, schema string) {
	dialect = strings.ToLower(c.Database.Drive)
	schema = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Db, c.Database.Charset)
	return
}

// XML 加载配置文件
func XML(env string) (data *Env, err error) {
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
func JSON(env string) (data *Env, err error) {
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
