package resource

import (
	`encoding/xml`
	
	`github.com/chaodoing/providers/console/environment/resource/asset`
	`github.com/chaodoing/providers/console/environment/resource/template`
	`github.com/chaodoing/providers/console/environment/resource/upload`
)

// Resource 资源配置
type Resource struct {
	XMLName  xml.Name           `xml:"resource"`
	Comment  string             `xml:"comment,attr"`
	Favicon  *Favicon           `xml:"favicon"`
	Template *template.Template `xml:"template"`
	Asset    *asset.Asset       `xml:"asset"`
	Upload   *upload.Upload     `xml:"upload"`
}

func NewResource() (data *Resource, err error) {
	var (
		favicon *Favicon
		tx      *template.Template
		at      *asset.Asset
		ud      *upload.Upload
	)
	favicon, err = NewFavicon()
	if err != nil {
		return
	}
	tx, err = template.NewTemplate()
	if err != nil {
		return
	}
	at, err = asset.NewAsset()
	if err != nil {
		return
	}
	ud, err = upload.NewUpload()
	if err != nil {
		return
	}
	data = &Resource{
		Comment:  "资源数据配置",
		Favicon:  favicon,
		Template: tx,
		Asset:    at,
		Upload:   ud,
	}
	return
}
