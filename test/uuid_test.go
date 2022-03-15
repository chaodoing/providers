package test

import (
	`crypto/md5`
	`encoding/base64`
	`fmt`
	`testing`
)

func TestUUID(t *testing.T) {
	// nano := strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10)
	encrypt := md5.Sum([]byte("123.com"))
	guid := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", encrypt)))
	t.Log(fmt.Sprintf("%x", encrypt))
	t.Log(guid)
	t.Log(len(guid))
}
