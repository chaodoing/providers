package test

import (
	`encoding/xml`
	`testing`
	
	`github.com/chaodoing/providers/putil`
)

type (
	host struct {
		XMLName  xml.Name `xml:"host"`
		XMLValue string   `xml:",innerxml"`
		State    string   `xml:"state,attr"`
	}
	port struct {
		XMLName  xml.Name `xml:"port"`
		XMLValue string   `xml:",innerxml"`
		State    string   `xml:"state,attr"`
	}
	cors struct {
		XMLName  xml.Name `xml:"cors"`
		XMLValue string   `xml:",innerxml"`
		State    string   `xml:"state,attr"`
	}
	listener struct {
		XMLName xml.Name `xml:"listener"`
		State   string   `xml:"state,attr"`
		Host    host     `xml:"host"`
		Port    port     `xml:"port"`
		Cors    cors     `xml:"cors"`
	}
	config struct {
		XMLName  xml.Name `xml:"config"`
		Listener listener `xml:"listener"`
	}
)

func TestSaveXML(t *testing.T) {
	var data = config{
		Listener: listener{
			State: "服务监听",
			Host: host{
				State:    "监听主机",
				XMLValue: "0.0.0.0",
			},
			Port: port{
				XMLValue: "80",
				State:    "监听端口",
			},
			Cors: cors{
				XMLValue: "false",
				State:    "允许跨域",
			},
		},
	}
	err := putil.SaveXML(data, "./index.xml")
	if err != nil {
		t.Error(err)
	}
	t.Log("Success")
}
