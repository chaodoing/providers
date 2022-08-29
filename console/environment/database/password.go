package database

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Password 数据库连接密码
type Password struct {
	XMLName  xml.Name `xml:"password"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证IP地址
func (h *Password) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("数据库连接密码不能为空")
	}
	return nil
}

// NewPassword 生成数据库连接密码
func NewPassword() (data *Password, err error) {
	data = &Password{
		Comment: "数据库连接密码",
	}
	command := promptui.Prompt{
		Label:    "数据库连接密码",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
