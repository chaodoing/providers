package upload

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Url 访问地址
type Url struct {
	XMLName  xml.Name `xml:"url"`
	Comment  string   `xml:"comment,attr"`
	XMLValue string   `xml:",innerxml"`
}

// validate 验证网站图标
func (h *Url) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("上传目录访问地址不能为空")
	}
	return nil
}

func NewUrl() (data *Url, err error) {
	data = &Url{
		Comment: "上传目录访问地址",
	}
	command := promptui.Prompt{
		Label:    "请输入上传目录访问地址",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
