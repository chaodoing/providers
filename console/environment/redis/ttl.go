package redis

import (
	`encoding/xml`
	`errors`
	`strconv`
	`strings`
	
	`github.com/manifoldco/promptui`
)

type TTL struct {
	XMLName  xml.Name `xml:"ttl"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证端口信息
func (h *TTL) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("Redis缓存时间不能为空")
	}
	
	_, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		return errors.New("Redis缓存时间必须是数值哦")
	}
	return
}

// NewTTL 缓存时间
func NewTTL() (data *TTL, err error) {
	data = &TTL{
		Comment: "Redis默认缓存时间",
	}
	command := promptui.Prompt{
		Label:    "Redis默认缓存时间",
		Validate: data.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
