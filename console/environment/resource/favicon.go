package resource

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Favicon 网站图标
type Favicon struct {
	XMLName  xml.Name `xml:"favicon"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证网站图标
func (h *Favicon) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("网站图标资源位置不能为空")
	}
	return nil
}

// NewFavicon 初始化网站图标位置
func NewFavicon() (data *Favicon, err error) {
	data = &Favicon{
		Comment: "网站图标资源位置",
	}
	command := promptui.Prompt{
		Label:    "请输入网站图标资源位置",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
