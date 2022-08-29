package encrypt

import (
	`strings`
)

// Password 加密密码
//  @param password string 要加密的密码
//  @return passwords string 加密后的密码
func Password(password string) (passwords string) {
	passwords = strings.ToUpper(Md5(password + password + password))
	return
}
