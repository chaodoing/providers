package util

import (
	`regexp`
	
	`github.com/gookit/validate`
)

func validatePassword(ps string) bool {
	if len(ps) < 6 || len(ps) > 21 {
		return false
	}
	count := 1
	num := `[0-9]{1}`
	aZ := `[a-z]{1}`
	AZ := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+-=|_.]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err == nil {
		count = count + 1
	}
	if b, err := regexp.MatchString(aZ, ps); !b || err == nil {
		
		count = count + 1
	}
	if b, err := regexp.MatchString(AZ, ps); !b || err == nil {
		count = count + 1
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err == nil {
		count = count + 1
	}
	return count >= 4
}

// validateConfirmPassword 确认密码
func validateConfirmPassword(valid *validate.Validation) interface{} {
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

// Validate 数据验证
//  @param data      要验证的结构体
//  @param scene     验证场景
//  @return has      是否有错误
//  @return messages 错误消息
func Validate(data interface{}, scene string, Scenarios validate.SValues) (hasErr bool, messages validate.Errors) {
	valid := validate.Struct(data)
	valid.AddValidator("password", validatePassword)
	valid.AddValidator("confirm_password", validateConfirmPassword(valid))
	valid.WithScenarios(Scenarios)
	return !valid.Validate(scene), valid.Errors
}
