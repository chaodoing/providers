package test

import (
	`testing`
	
	`github.com/chaodoing/providers/containers`
	`github.com/chaodoing/providers/util`
)

func TestReadXML(t *testing.T) {
	var data containers.Config
	
	if err := util.ReadXML("index.xml", &data); err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}