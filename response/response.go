package response

import (
	`bytes`
	`encoding/json`
	`html/template`
	
	`github.com/kataras/iris/v12`
	
	`github.com/chaodoing/providers/asset`
)

type Response struct {
	ctx  iris.Context
	data interface{}
}

// JsonHtml 数据转换为HTML内容
func (r Response) JsonHtml(title string) (content string, err error) {
	var bit []byte
	bit, err = asset.Asset("html/json.html")
	if err != nil {
		return
	}
	
	tpl, err := template.New("json").Parse(string(bit))
	if err != nil {
		return
	}
	bit, err = json.Marshal(r.data)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, map[string]string{
		"Title": title,
		"Json":  string(bit),
	})
	return buf.String(), err
}

func (r Response) Xml() error {
	r.ctx.Gzip(true)
	_, err := r.ctx.XML(r.data, iris.XML{Indent: "\t", Prefix: ""})
	return err
}

func (r Response) Json() error {
	r.ctx.Gzip(true)
	_, err := r.ctx.JSON(r.data, iris.JSON{Indent: "\t", Prefix: ""})
	return err
}

func (r Response) Send() (err error) {
	html, err := r.JsonHtml("JSON")
	if err != nil {
		return err
	}
	r.ctx.Gzip(true)
	r.ctx.Negotiation().JSON(r.data).XML(r.data).HTML(html).EncodingGzip()
	_, err = r.ctx.Negotiate(nil)
	return
}

// Success 输出成功内容
func (r Response) Success() (err error) {
	r.data = Data{
		Status:  0,
		Message: "OK",
		Data:    r.data,
	}
	return r.Send()
}

// Error 输出错误内容
func (r Response) Error(status uint32, message string) (err error) {
	r.data = Data{
		Status:  status,
		Message: message,
		Data:    r.data,
	}
	return r.Send()
}

// Pagination 分页内容输出
func (r Response) Pagination(data Pagination) (err error) {
	r.data = data
	return r.Send()
}

// Responsive 响应器实例化
func Responsive(ctx iris.Context, data interface{}) (response Response) {
	return Response{
		ctx:  ctx,
		data: data,
	}
}
