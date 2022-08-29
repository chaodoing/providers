package iris

import (
	`encoding/xml`
)

// Iris 框架日志
type Iris struct {
	XMLName xml.Name `xml:"iris"`
	Comment string   `xml:"comment,attr"`
	Level   *Level   `xml:"level"`
	File    *File    `xml:"file"`
	Error   *Error   `xml:"error"`
}

// NewIris 初始化iris框架日志配置
func NewIris() (iris *Iris, err error) {
	var (
		level *Level
		file  *File
		erro  *Error
	)
	level, err = NewLevel()
	if err != nil {
		return
	}
	file, err = NewFile()
	if err != nil {
		return
	}
	erro, err = NewError()
	if err != nil {
		return
	}
	iris = &Iris{
		Comment: "框架日志配置",
		Level:   level,
		File:    file,
		Error:   erro,
	}
	return
}
