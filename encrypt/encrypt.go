package encrypt

import (
	`crypto/md5`
	`encoding/base64`
	`fmt`
	`github.com/google/uuid`
	`net/url`
)

// Md5 加密字符串
//  @param value string 要加密的字符串
//  @return string  加密后的字符串
func Md5(value string) string {
	h := md5.Sum([]byte(value))
	md5String := fmt.Sprintf("%x", h)
	return md5String
}

// Password 加密密码
//  @param password string 要加密的密码
//  @return passwords string 加密后的密码
func Password(password string) (passwords string) {
	passwords = Md5(url.QueryEscape(base64.StdEncoding.EncodeToString([]byte(password))))
	return
}

// UUID 生成UUID
//  @return string uuid 字符串
func UUID() string {
	return uuid.New().String()
}
