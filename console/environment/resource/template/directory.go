package template

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
		return errors.New("模板目录位置不能为空")
	}
	return nil
}

// NewDirectory 生成Directory
func NewDirectory() (data *Directory, err error) {
	data = &Directory{
		Comment: "模板存储目录",
	}
	command := promptui.Prompt{
		Label:    "请输入模板存储目录",
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
