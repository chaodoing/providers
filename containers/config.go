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
	Location string   `json:"-" xml:"-"`
	App      struct {
		Host         string `json:"host" xml:"host"`
		Port         uint32 `json:"port" xml:"port"`
		AllowedCross bool   `json:"allowed_cross" xml:"allowed-cross"`
	} `json:"app" xml:"app"`
	
	Log struct {
		Console  bool   `json:"console" xml:"console"`
		OnSave   bool   `json:"on_save" xml:"on-save"`
		Level    string `json:"level" xml:"level"`
		Db       bool   `json:"db" xml:"db"`
		DbLevel  string `json:"db_level" xml:"db-level"`
		DbFile   string `json:"db_file" xml:"db-file"`
		IrisFile string `json:"iris_file" xml:"iris-file"`
		ErrFile  string `json:"err_file" xml:"err-file"`
	} `json:"log" xml:"log"`
	
	Database struct {
		Drive    string `json:"drive" xml:"drive"`
		Host     string `json:"host" xml:"host"`
		Port     uint32 `json:"port" xml:"port"`
		Dbname   string `json:"dbname" xml:"dbname"`
		Username string `json:"username" xml:"username"`
		Password string `json:"password" xml:"password"`
		Charset  string `json:"charset" xml:"charset"`
	} `json:"database" xml:"database"`
	
	Redis struct {
		Host    string `json:"host" xml:"host"`
		Port    uint32 `json:"port" xml:"port"`
		Auth    string `json:"auth" xml:"auth"`
		DbIndex uint8  `json:"db_index" xml:"db-index"`
		Expire  uint   `json:"expire" xml:"expire"`
	} `json:"redis" xml:"redis"`
	
	AssetBundle struct {
		Template      string `json:"template" xml:"template"`
		Favicon       string `json:"favicon" xml:"favicon"`
		StaticUri     string `json:"static_uri" xml:"static-uri"`
		Static        string `json:"static" xml:"static"`
		UploadUri     string `json:"upload_uri" xml:"upload-uri"`
		Upload        string `json:"upload" xml:"upload"`
		UploadMaximum uint   `json:"upload_maximum" xml:"upload-maximum"`
	} `json:"asset_bundle" xml:"asset-bundle"`
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
	return
}