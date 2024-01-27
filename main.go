package main

import (
	message "animatic-v2/Message"
	network "animatic-v2/Network"
	"animatic-v2/Structures"
	tui "animatic-v2/Tui"
	utils "animatic-v2/Utils"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGTERM:
			os.Exit(0)
		case syscall.SIGKILL:
			os.Exit(1)
		}
	}()

	// If hasn't internet connection, break the program
	if network.HasNetwork() == false {
		message.ErrorMessage("Hasn't internet connection")
		return
	}

	var destPath string
	animeName := tui.GetAnimeName()

	if animeName == "" {
		message.ErrorMessage("We didn't receive the name of the anime")
		return
	}

	animeSearcher := network.AnimeSearcher{AnimeName: animeName}
	animeSelectedName, animeURL := animeSearcher.SearchAnimeInAnimeFire()

	if runtime.GOOS != "windows" {
		destPath = filepath.Join("/chromeMedia/Series/", utils.SplitAnimeName(animeSelectedName))
	} else {
		userProfile := os.Getenv("USERPROFILE")
		destPath = filepath.Join(userProfile, "chromeMedia\\Series\\", utils.SplitAnimeName(animeSelectedName))
	}

	anime := network.Anime{URL: animeURL}
	epList := anime.GetEpisodes()

	downloader := network.Downloader{DestPath: destPath, Anime: Structures.Anime{Name: animeSelectedName, Url: animeURL}, EpList: epList}
	downloader.DownloadAll()
}
