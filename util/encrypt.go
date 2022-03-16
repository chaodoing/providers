package util

import (
	`crypto/md5`
	`encoding/base64`
	`fmt`
	`os`
	`strings`
)

// Md5 加密字符串
//  @param value string 要加密的字符串
//  @return string  加密后的字符串
func Md5(value string) string {
	h := md5.Sum([]byte(value))
	md5String := fmt.Sprintf("%x", h)
	return md5String
}

// UUID 生成UUID
//  @return string uuid 字符串
func UUID() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// Password 加密密码
//  @param password string 要加密的密码
//  @return passwords string 加密后的密码
func Password(password string) (passwords string) {
	passwords = strings.ToUpper(Md5(password + password + password))
	return
}

// PasswordBase64 加密密码
//  @param password string 要加密的密码
//  @return passwords string 加密后的密码
func PasswordBase64(password string) (base64password string) {
	return base64.StdEncoding.EncodeToString([]byte(Password(password)))
}