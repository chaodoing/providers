package listener

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
	
	`github.com/chaodoing/providers/util`
)

type Host struct {
	XMLName  xml.Name `xml:"host"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证IP地址
func (h *Host) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("监听的IP不能为空")
	}
	if !util.ISIPv4(input) {
		return errors.New("监听的IP地址格式不正确")
	}
	return nil
}

// NewHost 生成Host
func NewHost() (host *Host, err error) {
	host = &Host{
		Comment: "网站监听IP",
	}
	command := promptui.Prompt{
		Label:    "请输入网站监听IP",
		Validate: host.validate,
	}
	hostValue, err := command.Run()
	if err != nil {
		return nil, err
	}
	host.XMLValue = hostValue
	return
}
