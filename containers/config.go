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
		Host         string `json:"host" xml:"host" name:"监听主机"`
		Port         string `json:"port" xml:"port" name:"监听端口"`
		AllowedCross bool   `json:"allowed_cross" xml:"allowed-cross" name:"运行跨域"`
	} `json:"app" xml:"app" name:"站点名称"`
	
	Log struct {
		OutputToConsole bool   `json:"output_to_console" xml:"output-to-console" name:"日志是否输出到控制台"`
		SaveToFile      bool   `json:"save_to_file" xml:"save-to-file" name:"是否存储日志"`
		IrisLevel       string `json:"iris_level" xml:"iris-level" name:"iris日志等级"`
		GormOn          bool   `json:"gorm_on" xml:"gorm-on" name:"数据库日志是否开启"`
		GormLevel       string `json:"gorm_level" xml:"gorm-level" name:"数据库日志文件等级"`
		DbFile          string `json:"db_file" xml:"db-file" name:"数据库日志文件"`
		IrisFile        string `json:"iris_file" xml:"iris-file" name:"iris框架文件位置"`
		ErrFile         string `json:"err_file" xml:"err-file" name:"日志错误文件配置"`
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
		Template      string `json:"template" xml:"template" name:"模板目录"`
		Favicon       string `json:"favicon" xml:"favicon" name:"网站图标"`
		StaticUri     string `json:"static_uri" xml:"static-uri" name:"HTTP访问位置"`
		Static        string `json:"static" xml:"static" name:"静态资源绝对路径"`
		UploadUri     string `json:"upload_uri" xml:"upload-uri" name:"上传目录HTTP位置"`
		Upload        string `json:"upload" xml:"upload" name:"上传目录资源绝对路径"`
		UploadMaximum uint   `json:"upload_maximum" xml:"upload-maximum" name:"允许上传的最大文件大小"`
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
