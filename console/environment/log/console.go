package log

import (
	`encoding/xml`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Console 是否允许跨域
type Console struct {
	XMLName  xml.Name `xml:"console"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// NewConsole 初始化是否允许跨域
func NewConsole() (console *Console, err error) {
	console = &Console{
		Comment: "是否输出到控制台",
	}
	command := promptui.Select{
		Label: "请选择是否输出到控制台",
		Items: []string{"是", "否"},
	}
	var consoleValue string
	_, consoleValue, err = command.Run()
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(consoleValue, "是") {
		console.XMLValue = "true"
	} else {
		console.XMLValue = "false"
	}
	return
}
