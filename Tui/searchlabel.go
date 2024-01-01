package tui

import (
	message "animatic-v2/Message"

	"github.com/manifoldco/promptui"
)

func GetAnimeName() string {
	label := promptui.Prompt{
		Label: "Enter the name of the anime to be downloaded",
	}

	animeName, err := label.Run()

	if err != nil {
		message.ErrorMessage(err.Error())
		return ""
	}

	return animeName
}
