package template

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Extension 模板文件扩展名称
type Extension struct {
	XMLName  xml.Name `xml:"extension"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 检查目录
func (d *Extension) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("模板文件扩展名称不能为空")
	}
	return nil
}

// NewExtension 模板文件扩展名称
func NewExtension() (data *Extension, err error) {
	data = &Extension{
		Comment: "模板文件扩展名称",
	}
	command := promptui.Prompt{
		Label:    "请输入模板文件扩展名称",
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
