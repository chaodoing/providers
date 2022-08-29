package test

import (
	`fmt`
	`testing`
	
	`github.com/manifoldco/promptui`
)

func TestPrompt(t *testing.T) {
	prompt := promptui.Prompt{
		Label: "请输入应用描述",
	}
	result, err := prompt.Run()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result)
}
