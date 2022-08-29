package test

import (
	`encoding/json`
	`testing`
	
	`github.com/go-ini/ini`
	
	`github.com/chaodoing/providers/containers`
)

func TestINI(t *testing.T) {
	config, err := ini.Load("../.env")
	if err != nil {
		t.Error(err)
		return
	}
	var data = &containers.Env{}
	err = config.MapTo(data)
	if err != nil {
		t.Error(err)
		return
	}
	bit, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bit))
}
