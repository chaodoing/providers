package listener

import (
	`encoding/xml`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// Cors 是否允许跨域
type Cors struct {
	XMLName  xml.Name `xml:"cors"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// NewCors 初始化是否允许跨域
func NewCors() (cors *Cors, err error) {
	cors = &Cors{
		Comment: "是否允许跨域请求",
	}
	command := promptui.Select{
		Label: "请选择是否允许跨域请求",
		Items: []string{"是", "否"},
	}
	var crosValue string
	_, crosValue, err = command.Run()
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(crosValue, "是") {
		cors.XMLValue = "true"
	} else {
		cors.XMLValue = "false"
	}
	return
}
