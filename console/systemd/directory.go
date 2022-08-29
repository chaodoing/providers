package systemd

import (
	`errors`
	`strings`
	
	`github.com/gookit/goutil/fsutil`
	`github.com/manifoldco/promptui`
)

type directory string

func (d directory) validate(input string) (err error) {
	if strings.EqualFold(input, "") {
		return errors.New("服务工作路径不能为空")
	}
	if !fsutil.DirExist(input) {
		return errors.New("服务工作路径不存在")
	}
	return
}

func Directory() (data directory, err error) {
	command := promptui.Prompt{
		Label:    "应用工作路径",
		Validate: data.validate,
	}
	var value string
	value, err = command.Run()
	data = directory(value)
	return
}
