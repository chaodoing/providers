package test

import (
	`testing`
	
	`providers/containers`
	`providers/util`
)

func TestSaveJSON(t *testing.T) {
	var data containers.Config
	if err := util.ReadXML("index.xml", &data); err != nil {
		t.Error(err)
		return
	}
	if err := util.SaveJSON(data, "index.json"); err != nil {
		t.Error(err)
		return
	}
	t.Log("数据写入成功")
}

