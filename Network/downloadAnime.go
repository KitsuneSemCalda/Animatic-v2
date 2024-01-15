package network

import (
	message "animatic-v2/Message"
	"animatic-v2/Structures"
	utils "animatic-v2/Utils"
	"fmt"
	"path/filepath"
)

func DownloadAll(destPath string, anime Structures.Anime, epList []Structures.Episode) {
	for i := range epList {
		episode := utils.EpisodeFormatter(i + 1)
		newdestPath := filepath.Join(destPath, episode)

		if CheckWebsite(epList[i].URL) == false {
			message.ErrorMessage("Cant acess episode url")
			return
		}

		fmt.Println(newdestPath)

		videoUrl := extractVideoUrl(epList[i].URL)

		videoSrc := extractActualVideoLabel(videoUrl)

		downloadVideo(newdestPath, videoSrc)
	}
}
