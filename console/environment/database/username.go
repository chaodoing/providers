package database

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Username 数据库连接用户名
type Username struct {
	XMLName  xml.Name `xml:"username"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证IP地址
func (h *Username) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("数据库连接IP不能为空")
	}
	return nil
}

// NewUsername 生成数据库连接用户
func NewUsername() (data *Username, err error) {
	data = &Username{
		Comment: "数据库连接用户",
	}
	command := promptui.Prompt{
		Label:    "数据库连接用户",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
