package test

import (
	`testing`
	
	`github.com/gookit/goutil/envutil`
)

func TestEnv(t *testing.T) {
	t.Log(envutil.Getenv("ENV", "development"))
}
