package listener

import (
	`encoding/xml`
)

// Listener 监听内容
type Listener struct {
	XMLName xml.Name `xml:"listener"`
	Comment string   `xml:"comment,attr"`
	Host    *Host    `xml:"host"`
	Port    *Port    `xml:"port"`
	Cors    *Cors    `xml:"cors"`
}

// NewListener 创建监听信息
func NewListener() (listener *Listener, err error) {
	listener = &Listener{
		Comment: "服务监听",
	}
	var (
		host *Host
		port *Port
		cors *Cors
	)
	host, err = NewHost()
	if err != nil {
		return
	}
	port, err = NewPort()
	if err != nil {
		return
	}
	cors, err = NewCors()
	if err != nil {
		return
	}
	listener.Host = host
	listener.Port = port
	listener.Cors = cors
	return
}
