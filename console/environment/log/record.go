package log

import (
	`encoding/xml`
	`strings`
	
	`github.com/manifoldco/promptui`
)

type Record struct {
	XMLName  xml.Name `xml:"record"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// NewRecord 初始化是否记录到日志文件
func NewRecord() (console *Record, err error) {
	console = &Record{
		Comment: "是否记录到日志文件",
	}
	command := promptui.Select{
		Label: "请选择是否记录到日志文件",
		Items: []string{"是", "否"},
	}
	var value string
	_, value, err = command.Run()
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(value, "是") {
		console.XMLValue = "true"
	} else {
		console.XMLValue = "false"
	}
	return
}
