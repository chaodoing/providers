package encrypt

import (
	`crypto/md5`
	`fmt`
)

// Md5 加密字符串
//  @param value string 要加密的字符串
//  @return string  加密后的字符串
func Md5(value string) string {
	h := md5.Sum([]byte(value))
	md5String := fmt.Sprintf("%x", h)
	return md5String
}
