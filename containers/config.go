package containers

import (
	`encoding/json`
	`encoding/xml`
	`fmt`
	`os`
	`strings`
)

// Config 配置文件
type Config struct {
	XMLName  xml.Name `json:"-" xml:"root"`
	Location string   `json:"-" xml:"location,attr" name:"本地存储位置"`
	Version  string   `json:"-" xml:"version,attr"`
	App      struct {
		Host  string `json:"host" xml:"host" name:"监听主机"`
		Port  string `json:"port" xml:"port" name:"监听端口"`
		Cross bool   `json:"cross" xml:"cross" name:"运行跨域"`
	} `json:"app" xml:"app" name:"站点名称"`
	
	Log struct {
		Console bool `json:"console" xml:"console" name:"日志是否输出到控制台"`
		IsSave  bool `json:"is_save" xml:"is-save" name:"是否存储日志"`
		Iris    struct {
			Level   string `json:"level" xml:"level" name:"日志等级"`
			File    string `json:"file" xml:"file" name:"存储位置"`
			ErrFile string `json:"err_file" xml:"err-file" name:"错误文件"`
		} `json:"iris" xml:"iris"`
		Database struct {
			Enable bool   `json:"enable" xml:"enable" name:"开启日志"`
			Level  string `json:"level" xml:"level" name:"日志等级"`
			File   string `json:"file" xml:"file" name:"日志文件"`
		} `json:"database" xml:"database"`
	} `json:"log" xml:"log" name:"日志配置"`
	
	Database struct {
		Drive    string `json:"drive" xml:"drive" name:"日志驱动"`
		Host     string `json:"host" xml:"host" name:"连接主机"`
		Port     uint32 `json:"port" xml:"port" name:"连接端口"`
		Dbname   string `json:"dbname" xml:"dbname" name:"数据库名称"`
		Username string `json:"username" xml:"username" name:"连接用户"`
		Password string `json:"password" xml:"password" name:"连接密码"`
		Charset  string `json:"charset" xml:"charset" name:"数据库字符集"`
	} `json:"database" xml:"database" name:"数据库配置"`
	Redis struct {
		Host   string `json:"host" xml:"host" name:"连接主机"`
		Port   uint32 `json:"port" xml:"port" name:"连接端口"`
		Auth   string `json:"auth" xml:"auth" name:"认证密码"`
		Index  uint8  `json:"index" xml:"index" name:"数据库索引"`
		Expire uint   `json:"expire" xml:"expire" name:"redis存储TTL"`
	} `json:"redis" xml:"redis" name:"redis配置"`
	AssetBundle struct {
		Template string `json:"template" xml:"template" name:"模板目录"`
		Favicon  string `json:"favicon" xml:"favicon" name:"网站图标"`
		Static   struct {
			Uri  string `json:"uri" xml:"uri" name:"HTTP访问位置"`
			Path string `json:"path" xml:"path" name:"存储实际位置"`
		} `json:"static" xml:"static" name:"静态资源绝对路径"`
		Upload struct {
			Maximum uint   `json:"maximum" xml:"maximum" name:"允许上传的最大文件大小"`
			Uri     string `json:"uri" xml:"uri" name:"上传目录HTTP访问位置"`
			Path    string `json:"path" xml:"path" name:"上传目录存储实际位置"`
		} `json:"upload" xml:"upload"`
	} `json:"asset_bundle" xml:"asset-bundle" name:"静态资源配置"`
}

// dialect 转换数据配置
//	@return dialect=mysql
//	@return schema=root:123.com@tcp(127.0.0.1:3306)/arrangement?charset=utf8mb4&parseTime=True&loc=Local
//	@return logMode=false
func (c Config) dialect() (dialect, schema string) {
	dialect = strings.ToLower(c.Database.Drive)
	schema = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Dbname, c.Database.Charset)
	return
}

func ReadJSONConfig(file string) (data Config, err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(file))
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return
	}
	data.Version = os.Getenv("VERSION")
	data.Location = os.ExpandEnv(file)
	return
}
func ReadXMLConfig(file string) (data Config, err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(file))
	if err != nil {
		return
	}
	err = xml.Unmarshal(content, &data)
	if err != nil {
		return
	}
	data.Version = os.Getenv("VERSION")
	data.Location = os.ExpandEnv(file)
	return
}
