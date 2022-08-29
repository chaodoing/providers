package asset

import (
	`encoding/xml`
)

type Asset struct {
	XMLName   xml.Name   `xml:"asset"`
	Comment   string     `xml:"comment,attr"`
	Url       *Url       `xml:"url"`
	Directory *Directory `xml:"directory"`
}

func NewAsset() (data *Asset, err error) {
	var (
		url       *Url
		directory *Directory
	)
	url, err = NewUrl()
	if err != nil {
		return
	}
	directory, err = NewDirectory()
	if err != nil {
		return
	}
	data = &Asset{
		Comment:   "静态资源配置",
		Url:       url,
		Directory: directory,
	}
	return
}
