package database

import (
	`encoding/xml`
	
	`github.com/manifoldco/promptui`
)

type Drive struct {
	XMLName  xml.Name `xml:"drive"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// NewDrive 数控类型
func NewDrive() (data *Drive, err error) {
	data = &Drive{
		Comment: "数据库类型",
	}
	command := promptui.Select{
		Label: "请选择数据库类型",
		Items: []string{"mysql"},
	}
	_, value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
