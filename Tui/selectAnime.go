package tui

import (
	message "animatic-v2/Message"

	structure "animatic-v2/Structures"

	"github.com/manifoldco/promptui"
)

func SelectAnimes(animes []structure.Anime) int {

	/*
		animesNames := make([]string, 0)

		for i := range animes {
			animesNames := append(animesNames, animes[i].Name)
		}
	*/

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "▶ {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "▶ {{ .Name | cyan | underline }}",
	}

	prompt := promptui.Select{
		Label:     "Select the anime",
		Items:     animes,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		message.ErrorMessage(err.Error())
		return -1
	}

	return index
}
