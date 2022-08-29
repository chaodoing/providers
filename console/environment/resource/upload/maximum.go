package upload

import (
	`encoding/xml`
	`errors`
	`strconv`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Maximum 允许上传的文件大小
type Maximum struct {
	XMLName  xml.Name `xml:"maximum"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证端口信息
func (h *Maximum) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("允许上传的最大文件大小")
	}
	_, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		errors.New("允许上传的最大文件大小只能是数值")
	}
	return
}

// NewMaximum 允许上传的最大
func NewMaximum() (data *Maximum, err error) {
	data = &Maximum{
		Comment: "允许上传的最大文件大小",
	}
	command := promptui.Prompt{
		Label:    "允许上传的最大文件大小[MB]",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
