package test

import (
	`testing`
	
	`providers/containers`
	`providers/util`
)

func TestSaveXML(t *testing.T) {
	var data containers.Config
	if err := util.SaveXML(data, "index.xml"); err != nil {
		t.Error(err)
		return
	}
	t.Log("数据写入成功")
}

