package test

import (
	`testing`
	
	`github.com/chaodoing/providers/containers`
	`github.com/chaodoing/providers/util`
)

func TestSaveXML(t *testing.T) {
	var data containers.Config
	data.Version = "v1.0.0"
	if err := util.SaveXML(data, "index.xml"); err != nil {
		t.Error(err)
		return
	}
	t.Log("数据写入成功")
}

