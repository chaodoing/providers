package systemd

import (
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

type description string

func (d description) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("应用描述不能为空")
	}
	return
}

func Description() (data description, err error) {
	command := promptui.Prompt{
		Label:    "请输入应用描述",
		Validate: data.validate,
	}
	var value string
	value, err = command.Run()
	data = description(value)
	return
}
