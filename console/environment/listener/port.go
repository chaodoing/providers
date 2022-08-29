package listener

import (
	`encoding/xml`
	`errors`
	`strconv`
	`strings`
	
	`github.com/manifoldco/promptui`
)

type Port struct {
	XMLName  xml.Name `xml:"port"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证IP地址
func (h *Port) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("监听的端口不能为空")
	}
	var data int64
	data, err = strconv.ParseInt(input, 10, 64)
	if data > 65535 || data < 80 {
		return errors.New("监听的端口格式错误[80 ~ 65535]")
	}
	return
}

// NewPort 生成 Port
func NewPort() (host *Port, err error) {
	host = &Port{
		Comment: "网站监听端口",
	}
	command := promptui.Prompt{
		Label:    "请输入网站监听端口",
		Validate: host.validate,
	}
	portValue, err := command.Run()
	if err != nil {
		return nil, err
	}
	host.XMLValue = portValue
	return
}
