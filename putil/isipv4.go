package putil

import (
	`strconv`
	`strings`
)

// ISIPv4 是否是IPV4
func ISIPv4(IP string) bool {
	// 字符串这样切割
	strs := strings.Split(IP, ".")
	if len(strs) != 4 {
		return false
	}
	for _, s := range strs {
		if len(s) == 0 || (len(s) > 1 && s[0] == '0') {
			return false
		}
		// 直接访问字符串的值
		if s[0] < '0' || s[0] > '9' {
			return false
		}
		// 字符串转数字
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if n < 0 || n > 255 {
			return false
		}
	}
	return true
}
