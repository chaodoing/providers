package systemd

import (
	`errors`
	`strings`
	
	`github.com/manifoldco/promptui`
)

type execute string

func (d execute) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("应用执行命令不能为空")
	}
	return
}

func Execute() (data execute, err error) {
	command := promptui.Prompt{
		Label:    "应用执行命令",
		Validate: data.validate,
	}
	var value string
	value, err = command.Run()
	data = execute(value)
	return
}
