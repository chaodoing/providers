package iris

import (
	`encoding/xml`
	
	`github.com/manifoldco/promptui`
)

// Level 日志等级
type Level struct {
	XMLName  xml.Name `xml:"level"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
	Value    string   `xml:"value,attr"`
}

// NewLevel 初始化日志等级
func NewLevel() (level *Level, err error) {
	level = &Level{
		Comment: "iris日志等级",
		Value:   `disable, fatal, error, warn, info, debug`,
	}
	command := promptui.Select{
		Label: "请选择iris日志等级",
		Items: []string{"disable", "fatal", "error", "warn", "info", "debug"},
	}
	var value string
	_, value, err = command.Run()
	if err != nil {
		return nil, err
	}
	level.XMLValue = value
	return
}
