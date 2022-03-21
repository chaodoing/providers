package test

import (
	`testing`
	
	`github.com/chaodoing/providers/containers`
	`github.com/chaodoing/providers/util`
)

func TestGetTags(t *testing.T) {
	config := containers.Config{}
	t.Log(util.Field(config.App))
}
