package containers

import (
	`crypto/md5`
	`encoding/base64`
	`encoding/json`
	`errors`
	`fmt`
	`math/rand`
	`strconv`
	`strings`
	`time`
	
	`github.com/go-redis/redis`
	`github.com/kataras/iris/v12`
)

const (
	Basic  = "Basic "
	Bearer = "Bearer "
	Prefix = "MEMBERSTATE:"
)

type Authorized struct {
	rdx           *redis.Client
	Expire        int64  `json:"expire" xml:"expire"`
	Authorization string `json:"authorization" xml:"authorization"`
	Digest        string `json:"digest" xml:"digest"`
}

// Clear 用户缓存
func (a *Authorized) Clear(ctx iris.Context) error {
	token, err := a.authString(ctx)
	if err != nil {
		return err
	}
	if !strings.EqualFold(token, "") {
		index := Prefix + token
		_, err := a.rdx.Del(index).Result()
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("AccessToken Not Found")
}

// Deposit 存入数据
//  @return error err 错误内容
func (a *Authorized) Deposit(data interface{}) (err error) {
	var (
		b     []byte
		index string
	)
	b, err = json.Marshal(data)
	if err != nil {
		return err
	}
	if strings.EqualFold(a.Authorization, "") {
		a.Authorization = a.stalkString()
	}
	a.Digest = string(b)
	index = Prefix + a.Authorization
	if _, err = a.rdx.Set(index, string(b), time.Duration(a.Expire*int64(time.Second))).Result(); err != nil {
		return err
	}
	return
}

// Auth 用户信息认证
func (a *Authorized) Auth(ctx iris.Context, data interface{}) (err error) {
	var Authorization string
	Authorization, err = a.authString(ctx)
	if err != nil {
		return
	}
	key := Prefix + Authorization
	if a.Digest, err = a.rdx.Get(key).Result(); err != nil {
		return errors.New("用户登录信息找不到")
	}
	
	err = json.Unmarshal([]byte(a.Digest), &data)
	if err != nil {
		return err
	}
	
	if strings.EqualFold(ctx.GetHeader("Refresh-Token"), "false") || strings.EqualFold(ctx.GetHeader("Refresh-Token"), "0") {
		ctx.Header("Refresh-Token", "false")
		ctx.Header("Refresh-Expires", fmt.Sprintf("%d", time.Now().Unix()+a.Expire))
		if _, err = a.rdx.Set(key, a.Digest, time.Duration(a.Expire*int64(time.Second))).Result(); err != nil {
			return err
		}
	} else {
		if _, err = a.rdx.Del(Authorization).Result(); err != nil {
			return err
		}
		a.Authorization = a.stalkString()
		ctx.Header("Refresh-Token", a.Authorization)
		ctx.Header("Refresh-Expires", fmt.Sprintf("%d", time.Now().Unix()+a.Expire))
		index := Prefix + ":" + a.Authorization
		if _, err = a.rdx.Set(index, a.Digest, time.Duration(a.Expire*int64(time.Second))).Result(); err != nil {
			return err
		}
	}
	return
}

// stalkString 生成用户跟踪字符串
func (a *Authorized) stalkString() string {
	nano := strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10)
	encrypt := md5.Sum([]byte(nano))
	return strings.ToUpper(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", encrypt))))
}

// authString 获取用户认证字符串
func (a *Authorized) authString(ctx iris.Context) (string, error) {
	if token := ctx.GetHeader("Accept-Token"); !strings.EqualFold(token, "") {
		return token, nil
	}
	if token := ctx.GetHeader("Authorization"); !strings.EqualFold(token, "") {
		if token = strings.TrimPrefix(token, Basic); !strings.EqualFold(token, "") {
			return token, nil
		}
		if token = strings.TrimPrefix(token, Bearer); !strings.EqualFold(token, "") {
			return token, nil
		}
	}
	return "", errors.New("找不到用户认证字符串")
}
