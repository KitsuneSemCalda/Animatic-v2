package tui

import "github.com/manifoldco/promptui"

func SearchAnime() (string, error) {
	prompt := promptui.Prompt{
		Label: "Type the anime from download",
	}

	return prompt.Run()
}
