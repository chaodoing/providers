package template

import (
	`encoding/xml`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Reload 重新加载模板
type Reload struct {
	XMLName  xml.Name `xml:"reload"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// NewReload 是否重新加载模板
func NewReload() (data *Reload, err error) {
	data = &Reload{
		Comment: "是否每次访问都重新加载模板",
	}
	command := promptui.Select{
		Label: "是否每次访问都重新加载模板",
		Items: []string{"是", "否"},
	}
	var consoleValue string
	_, consoleValue, err = command.Run()
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(consoleValue, "是") {
		data.XMLValue = "true"
	} else {
		data.XMLValue = "false"
	}
	return
}
