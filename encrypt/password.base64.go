package encrypt

import (
	`encoding/base64`
)

// PasswordBase64 加密密码
//  @param password string 要加密的密码
//  @return passwords string 加密后的密码
func PasswordBase64(password string) (base64password string) {
	return base64.StdEncoding.EncodeToString([]byte(Password(password)))
}
