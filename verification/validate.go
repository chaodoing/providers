package verification

import (
	`github.com/gookit/validate`
)

// Validate 数据验证
//  @param data      要验证的结构体
//  @param scene     验证场景
//  @return has      是否有错误
//  @return messages 错误消息
func Validate(data interface{}, scene string, Scenarios validate.SValues) (hasErr bool, messages validate.Errors) {
	valid := validate.Struct(data)
	valid.AddValidator("password", Password)
	valid.AddValidator("confirm_password", ConfirmPassword(valid))
	valid.WithScenarios(Scenarios)
	return !valid.Validate(scene), valid.Errors
}
