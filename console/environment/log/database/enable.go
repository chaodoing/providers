package database

import (
	`encoding/xml`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Enable 是否启用日志
type Enable struct {
	XMLName  xml.Name `xml:"enable"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// NewEnable 初始化是否启用日志
func NewEnable() (data *Enable, err error) {
	data = &Enable{
		Comment: "启用数据库日志",
	}
	command := promptui.Select{
		Label: "请选择是否启用数据库日志",
		Items: []string{"是", "否"},
	}
	var value string
	_, value, err = command.Run()
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(value, "是") {
		data.XMLValue = "true"
	} else {
		data.XMLValue = "false"
	}
	return
}
