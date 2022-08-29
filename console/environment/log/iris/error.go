package iris

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

type Error struct {
	XMLName  xml.Name `xml:"error"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证IP地址
func (h *Error) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("请输入错误日志文件名称")
	}
	return nil
}

// NewError 生成Host
func NewError() (error *Error, err error) {
	error = &Error{
		Comment: "错误日志文件名称",
	}
	command := promptui.Prompt{
		Label:    "错误日志文件名称",
		Validate: error.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	error.XMLValue = value
	return
}
