package util

import (
	`bytes`
	`encoding/json`
	`html/template`
	
	`github.com/kataras/iris/v12`
	
	`github.com/chaodoing/providers/assets`
)

type Respond struct {
	ctx  iris.Context
	data map[string]interface{}
}

// Data 设置数据
func (r Respond) Data(data map[string]interface{}) Respond {
	r.data = data
	return r
}

func (r Respond) JsonView(data interface{}) (string, error) {
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

// Send 发送数据到浏览器
func (r Respond) Send() {
	html, err := r.JsonView(r.data)
	if err != nil {
		r.ctx.Application().Logger().Error(err)
	}
	r.ctx.Gzip(true)
	r.ctx.Negotiation().JSON(r.data).XML(r.data).HTML(html).EncodingGzip()
	if _, err := r.ctx.Negotiate(nil); err != nil {
		r.ctx.Application().Logger().Error(err)
	}
	return
}

// SendData 发送设置数据到浏览器
func (r Respond) SendData(data interface{}) {
	html, err := r.JsonView(data)
	if err != nil {
		r.ctx.Application().Logger().Error(err)
	}
	r.ctx.Gzip(true)
	r.ctx.Negotiation().JSON(data).XML(data).HTML(html).EncodingGzip()
	if _, err := r.ctx.Negotiate(nil); err != nil {
		r.ctx.Application().Logger().Error(err)
	}
	return
}

func (r Respond) O(errorCode uint16, message string, data interface{}) {
	r.data["error_code"] = errorCode
	r.data["message"] = message
	r.data["data"] = data
	r.Send()
	return
}

func (r Respond) Success(data interface{}, msg ...string) {
	var message = "OK"
	if len(msg) > 0 {
		message = msg[0]
	}
	r.data["error_code"] = 0
	r.data["message"] = message
	r.data["data"] = data
	r.Send()
	return
}

func (r Respond) Error(message string, data ...interface{}) {
	r.data["error_code"] = 1
	r.data["message"] = message
	if len(data) > 0 {
		r.data["data"] = data[0]
	}
	r.Send()
	return
}

func O(ctx iris.Context, errorCode uint16, message string, data interface{}) {
	respond := NewRespond(ctx)
	respond.O(errorCode, message, data)
	return
}

func NewRespond(ctx iris.Context) Respond {
	return Respond{ctx: ctx}
}
