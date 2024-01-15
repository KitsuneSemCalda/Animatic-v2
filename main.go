package main

import (
	message "animatic-v2/Message"
	network "animatic-v2/Network"
	"animatic-v2/Structures"
	tui "animatic-v2/Tui"
	utils "animatic-v2/Utils"
	"path/filepath"
)

func main() {

	// If hasn't internet connection, break the program
	internetConnection := network.HasNetwork()

	if internetConnection == false {
		message.ErrorMessage("Hasn't internet connection")
		return
	}
	/*
		 TODO: Create a logic to multiple website searching
		 TODO: Create a logic to unique list from website searching (sqlite3) keep the following infomation:
		 	- Name of Anime string
			- Website string
			- IsDownloadable boolean
			- NumberOfEpisode integer
	*/

	animeName := tui.GetAnimeName()

	if animeName == "" {
		message.ErrorMessage("We didn't receive the name of the anime")
		return
	}

	animeURL, animeSelectedName, err := network.SearchAnimeInAnimeFire(animeName)
	destPath := filepath.Join("/chromeMedia/Series/", utils.SplitAnimeName(animeSelectedName))

	if err != nil {
		message.ErrorMessage("Failed to Locate anime")
	}

	epList := network.GetAnimeEpisodes(animeURL)

	network.DownloadAll(destPath, Structures.Anime{Name: animeSelectedName, Url: animeURL}, epList)
}
