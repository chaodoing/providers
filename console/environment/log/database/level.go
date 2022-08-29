package database

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

// "silent": logger.Silent,
// "error":  logger.Error,
// "warn":   logger.Warn,
// "info":   logger.Info,

// NewLevel 初始化日志等级
func NewLevel() (data *Level, err error) {
	data = &Level{
		Comment: "gorm日志等级",
		Value:   `silent, error, warn, info`,
	}
	command := promptui.Select{
		Label: "请选择gorm日志等级",
		Items: []string{"silent", "error", "warn", "info"},
	}
	var value string
	_, value, err = command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
