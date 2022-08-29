package iris

import (
	`encoding/xml`
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

// File 日志文件名称
type File struct {
	XMLName  xml.Name `xml:"file"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// validate 验证IP地址
func (h *File) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("请输入iris日志文件名称")
	}
	return nil
}

// NewFile 生成Host
func NewFile() (error *File, err error) {
	error = &File{
		Comment: "iris日志文件名称",
	}
	command := promptui.Prompt{
		Label:    "iris日志文件名称",
		Validate: error.validate,
	}
	value, err := command.Run()
	if err != nil {
		return nil, err
	}
	error.XMLValue = value
	return
}
