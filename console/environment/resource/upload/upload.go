package upload

import (
	`encoding/xml`
)

type Upload struct {
	XMLName   xml.Name   `xml:"upload"`
	Comment   string     `xml:"comment,attr"`
	Maximum   *Maximum   `xml:"maximum"`
	Url       *Url       `xml:"url"`
	Directory *Directory `xml:"directory"`
}

func NewUpload() (data *Upload, err error) {
	var (
		directory *Directory
		url       *Url
		maximum   *Maximum
	)
	maximum, err = NewMaximum()
	if err != nil {
		return
	}
	url, err = NewUrl()
	if err != nil {
		return
	}
	directory, err = NewDirectory()
	if err != nil {
		return
	}
	data = &Upload{
		Comment:   "模板配置",
		Maximum:   maximum,
		Url:       url,
		Directory: directory,
	}
	return
}
