package util

import (
	`bytes`
	`encoding/json`
	`encoding/xml`
	`html/template`
	`strings`
	
	`github.com/kataras/iris/v12`
	`gorm.io/gorm`
	
	`github.com/chaodoing/providers/assets`
)

type (
	Data struct {
		XMLName xml.Name    `xml:"root" json:"-"`
		Status  uint16      `json:"error_code" xml:"error_code"`
		Message string      `json:"message" xml:"message"`
		Data    interface{} `json:"data" xml:"data"`
	}
	Page struct {
		XMLName xml.Name    `xml:"root" json:"-"`
		Status  uint16      `json:"error_code" xml:"error_code"`
		Message string      `json:"message" xml:"message"`
		Data    interface{} `json:"data" xml:"data"`
		Total   uint64      `json:"total" xml:"total"`
		Limit   uint        `json:"limit" xml:"limit"`
		Page    uint        `json:"page" xml:"page"`
	}
)


func jsonView(data interface{}) (string, error) {
	html, _ := assets.Asset("template/json.html")
	tpl, err := template.New("jsonView").Parse(string(html))
	if err != nil {
		return "", err
	}
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, map[string]string{
		"Title": "JsonView",
		"Json":  string(jsonByte),
	})
	return buf.String(), err
}

func (p Page) O(ctx iris.Context) {
	if p.Status == 0 {
		p.Message = strings.ToUpper(p.Message)
	}
	html, _ := jsonView(p)
	ctx.Gzip(true)
	ctx.Negotiation().JSON(p).XML(p).HTML(html).EncodingGzip()
	if _, err := ctx.Negotiate(nil); err != nil {
		ctx.Application().Logger().Error(err)
	}
}

func O(ctx iris.Context, status uint16, message string, datum interface{}) {
	var data = Data{
		Status:  status,
		Message: message,
		Data:    datum,
	}
	if data.Status == 0 {
		data.Message = strings.ToUpper(data.Message)
	}
	html, _ := jsonView(data)
	ctx.Gzip(true)
	
	ctx.Negotiation().EncodingGzip().JSON(data).XML(data).HTML(html)
	if _, err := ctx.Negotiate(nil); err != nil {
		ctx.Application().Logger().Error(err)
	}
	return
}

func Pagination(db *gorm.DB, table string, page, limit int) (data Page, offset int) {
	var (
		total int64
		err   error
	)
	if page > 0 {
		offset = (page - 1) * limit
	} else {
		page = 1
	}
	if err = db.Table(table).Count(&total).Error; err != nil {
		data = Page{
			Status:  1,
			Message: err.Error(),
			Page:    uint(page),
			Total:   uint64(total),
			Limit:   uint(limit),
			Data:    nil,
		}
		return
	}
	data = Page{
		Status:  0,
		Message: "OK",
		Page:    uint(page),
		Total:   uint64(total),
		Limit:   uint(limit),
		Data:    nil,
	}
	return
}