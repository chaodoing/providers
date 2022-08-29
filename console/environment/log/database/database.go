package database

import (
	`encoding/xml`
)

type Database struct {
	XMLName xml.Name `xml:"database"`
	Comment string   `xml:"comment,attr"`
	Enable  *Enable  `xml:"enable"`
	File    *File    `xml:"file"`
	Level   *Level   `xml:"level"`
}

// NewDatabase 初始化iris框架日志配置
func NewDatabase() (data *Database, err error) {
	var (
		level  *Level
		file   *File
		enable *Enable
	)
	level, err = NewLevel()
	if err != nil {
		return
	}
	file, err = NewFile()
	if err != nil {
		return
	}
	enable, err = NewEnable()
	if err != nil {
		return
	}
	data = &Database{
		Comment: "框架日志配置",
		Level:   level,
		File:    file,
		Enable:  enable,
	}
	return
}
