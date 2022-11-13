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
)

// Track 用户跟踪
type Track struct {
	redisCli *redis.Client
	ctx      iris.Context
	TTL      time.Duration // TTL 存储有效期
	Token    string        // Token 访问Token
	// refresh 用户刷新数据
	Refresh struct {
		Token  string    // Token 新用户认证字符串
		Expire time.Time // Expire 新认证字符串有效期
	}
}

// expire 获取有效时间
func (t Track) expire() (date time.Time) {
	date = time.Now().Add(t.TTL)
	return
}

// get 获取当前访问的token
func (t Track) get() (token string, err error) {
	if token := t.ctx.GetHeader("Accept-Token"); !strings.EqualFold(token, "") {
		return strings.TrimPrefix(strings.TrimPrefix(token, Basic), Bearer), nil
	}
	if token := t.ctx.GetHeader("Authorization"); !strings.EqualFold(token, "") {
		if token = strings.TrimPrefix(token, Basic); !strings.EqualFold(token, "") {
			return strings.TrimPrefix(strings.TrimPrefix(token, Basic), Bearer), nil
		}
		if token = strings.TrimPrefix(token, Bearer); !strings.EqualFold(token, "") {
			return strings.TrimPrefix(strings.TrimPrefix(token, Basic), Bearer), nil
		}
	}
	return "", errors.New("找不到用户认证字符串")
}

// token 生成用户认证字符串
func (t Track) token() (token string) {
	nano := strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10)
	encrypt := md5.Sum([]byte(nano))
	token = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", encrypt)))
	return
}

// Del 清理当前的token
func (t Track) Del(ctx iris.Context) (err error) {
	t.ctx = ctx
	t.Token, err = t.get()
	if err != nil {
		return
	}
	_, err = t.redisCli.Exists(t.Token).Result()
	if err != nil {
		return
	}
	_, err = t.redisCli.Del(t.Token).Result()
	return
}

// Create 创建用户信息
func (t Track) Create(ctx iris.Context, data interface{}) (err error) {
	t.ctx = ctx
	t.Token = t.token()
	t.ctx.Header("Access-Control-Allow-Headers", "Refresh-Token, Accept-Version, Authorization, Accept-Token, Language, Access-Control-Allow-Methods, Access-Control-Allow-Origin, Cache-Control, Content-Type, if-match, if-modified-since, if-none-match, if-unmodified-since, X-Requested-With")
	t.ctx.Header("Access-Control-Expose-Headers", "Authorization, Accept-Token, Refresh-Token, Refresh-Expires")
	t.ctx.Header("Refresh-Expires", t.expire().Format("2006-01-02 15:04:05"))
	t.ctx.Header("Refresh-Token", t.Token)
	
	var bit []byte
	bit, err = json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = t.redisCli.Set(t.Token, string(bit), t.TTL).Result()
	return
}

// Get 获取用户认证信息
func (t Track) Get(ctx iris.Context, data interface{}) (err error) {
	t.ctx = ctx
	t.Token, err = t.get()
	if err != nil {
		return
	}
	var (
		RefreshToken = ctx.GetHeader("Refresh-Token")
		content      string
	)
	content, err = t.redisCli.Get(t.Token).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		return
	}
	t.ctx.Header("Access-Control-Allow-Headers", "Refresh-Token, Accept-Version, Authorization, Accept-Token, Language, Access-Control-Allow-Methods, Access-Control-Allow-Origin, Cache-Control, Content-Type, if-match, if-modified-since, if-none-match, if-unmodified-since, X-Requested-With")
	t.ctx.Header("Access-Control-Expose-Headers", "Authorization, Accept-Token, Refresh-Token, Refresh-Expires")
	t.ctx.Header("Refresh-Expires", t.expire().Format("2006-01-02 15:04:05"))
	if strings.EqualFold(RefreshToken, "0") || strings.EqualFold(RefreshToken, "false") || strings.EqualFold(RefreshToken, "off") {
		t.ctx.Header("Refresh-Token", "false")
		_, err = t.redisCli.Set(t.Token, content, t.TTL).Result()
		if err != nil {
			return
		}
	} else {
		_, err = t.redisCli.Del(t.Token).Result()
		if err != nil {
			return
		}
		t.Token = t.token()
		t.ctx.Header("Refresh-Token", t.Token)
		_, err = t.redisCli.Set(t.Token, content, t.TTL).Result()
		if err != nil {
			return
		}
	}
	return
}

// Update 更新用户信息
func (t Track) Update(ctx iris.Context, data interface{}) (err error) {
	t.ctx = ctx
	var token = ctx.ResponseWriter().Header().Get("Refresh-Token")
	if strings.EqualFold(token, "false") {
		t.Token, err = t.get()
		if err != nil {
			return
		}
	} else {
		t.Token = token
	}
	var bit []byte
	bit, err = json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = t.redisCli.Set(t.Token, string(bit), t.TTL).Result()
	if err != nil {
		return
	}
	return
}
