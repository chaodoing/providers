package redis

import (
	`encoding/xml`
	`errors`
	`strconv`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Port 数据库连接端口
type Port struct {
	XMLName  xml.Name `xml:"port"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证IP地址
func (h *Port) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("Redis连接端口不能为空")
	}
	var data int64
	data, err = strconv.ParseInt(input, 10, 64)
	if data > 65535 || data < 80 {
		return errors.New("Redis连接端口格式错误[80 ~ 65535]")
	}
	return
}

// NewPort 生成 Port
func NewPort() (data *Port, err error) {
	data = &Port{
		Comment: "Redis连接端口",
	}
	command := promptui.Prompt{
		Label:    "Redis连接端口",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
