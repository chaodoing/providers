package database

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Db 数据库名称
type Db struct {
	XMLName  xml.Name `xml:"db"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证数据库名称
func (h *Db) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("数据库名称不能为空")
	}
	return nil
}

// NewDb 初始化数据库名称
func NewDb() (data *Db, err error) {
	data = &Db{
		Comment: "数据库名称",
	}
	command := promptui.Prompt{
		Label:    "请输入数据库名称",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
