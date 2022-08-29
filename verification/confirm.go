package verification

import (
	`github.com/gookit/validate`
)

// ConfirmPassword 确认密码
func ConfirmPassword(valid *validate.Validation) interface{} {
	return func(value, attribute string) bool {
		if password, has := valid.Get(attribute); has {
			if pwd, ok := password.(string); ok {
				return pwd == value
			} else {
				return false
			}
		}
		return true
	}
}
