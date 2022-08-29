package template

import (
	`encoding/xml`
)

// Template 模板配置
type Template struct {
	XMLName   xml.Name   `xml:"template"`
	Comment   string     `xml:"comment,attr"`
	Directory *Directory `xml:"directory"`
	Extension *Extension `xml:"extension"`
	Reload    *Reload    `xml:"reload"`
}

func NewTemplate() (data *Template, err error) {
	var (
		directory *Directory
		extension *Extension
		reload    *Reload
	)
	directory, err = NewDirectory()
	if err != nil {
		return
	}
	extension, err = NewExtension()
	if err != nil {
		return
	}
	reload, err = NewReload()
	if err != nil {
		return
	}
	data = &Template{
		Comment:   "模板配置",
		Directory: directory,
		Extension: extension,
		Reload:    reload,
	}
	return
}
