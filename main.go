package main

import (
	message "animatic-v2/Message"
	network "animatic-v2/Network"
	tui "animatic-v2/Tui"
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

	network.SearchAnimeInAnimeFire(animeName)
}
