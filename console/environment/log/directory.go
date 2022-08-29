package log

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
		return errors.New("日志目录不能为空")
	}
	return nil
}

// NewDirectory 生成Directory
func NewDirectory() (directory *Directory, err error) {
	directory = &Directory{
		Comment: "日志存储目录",
	}
	command := promptui.Prompt{
		Label:    "请输入日志存储目录",
		Validate: directory.validate,
	}
	var value string
	value, err = command.Run()
	if err != nil {
		return nil, err
	}
	directory.XMLValue = value
	return
}
