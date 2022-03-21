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
	Location string   `json:"-" xml:"-" name:"本地存储位置"`
	Version  string  `json:"-" xml:"version,attr"`
	App      struct {
		Host         string `json:"host" xml:"host" name:"监听主机"`
		Port         string `json:"port" xml:"port" name:"监听端口"`
		AllowedCross bool   `json:"allowed_cross" xml:"allowed-cross" name:"运行跨域"`
	} `json:"app" xml:"app" name:"站点名称"`
	
	Wechat struct {
		MiniProgram struct {
			AppID     string `json:"app_id" xml:"app-id" name:"微信小程序APP_ID"`
			AppSecret string `json:"app_secret" xml:"app-secret" name:"微信小程序App_Secret"`
		} `json:"mini_program" xml:"mini-program" name:"微信小程序配置"`
		OfficialAccount struct {
			AppID          string `json:"app_id" xml:"app-id" name:"微信公众号APP_ID"`
			AppSecret      string `json:"app_secret" xml:"app-secret" name:"微信公众号App_Secret"`
			Token          string `json:"token" xml:"token" name:"验证令牌"`
			EncodingAESKey string `json:"encoding_aes_key" xml:"encoding-aes-key" name:"开启AES加密"`
		} `json:"official_account" xml:"official-account" name:"微信公众号配置"`
	} `json:"wechat" xml:"wechat" name:"微信数据配置"`
	
	Log struct {
		Console  bool   `json:"console" xml:"console" name:"日志是否开启"`
		OnSave   bool   `json:"on_save" xml:"on-save" name:"是否存储日志"`
		Level    string `json:"level" xml:"level" name:"iris日志等级"`
		Db       bool   `json:"db" xml:"db" name:"数据库日志是否开启"`
		DbLevel  string `json:"db_level" xml:"db-level" name:"数据库日志文件等级"`
		DbFile   string `json:"db_file" xml:"db-file" name:"数据库日志文件"`
		IrisFile string `json:"iris_file" xml:"iris-file" name:"iris框架文件位置"`
		ErrFile  string `json:"err_file" xml:"err-file" name:"日志错误文件配置"`
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
		Host    string `json:"host" xml:"host" name:"连接主机"`
		Port    uint32 `json:"port" xml:"port" name:"连接端口"`
		Auth    string `json:"auth" xml:"auth" name:"认证密码"`
		DbIndex uint8  `json:"db_index" xml:"db-index" name:"数据库索引"`
		Expire  uint   `json:"expire" xml:"expire" name:"redis存储TTL"`
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
	return
}
