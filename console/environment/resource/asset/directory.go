package asset

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

type Directory struct {
	XMLName  xml.Name `xml:"directory"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 检查目录
func (d *Directory) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("静态资源实际位置不能为空")
	}
	return nil
}

// NewDirectory 生成Directory
func NewDirectory() (data *Directory, err error) {
	data = &Directory{
		Comment: "静态资源实际位置",
	}
	command := promptui.Prompt{
		Label:    "请输入静态资源实际位置",
		Validate: data.validate,
	}
	var value string
	value, err = command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
