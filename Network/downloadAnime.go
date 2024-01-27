package network

import (
	message "animatic-v2/Message"
	"animatic-v2/Structures"
	utils "animatic-v2/Utils"
	"path/filepath"
)

type Downloader struct {
	DestPath string
	Anime    Structures.Anime
	EpList   []Structures.Episode
}

func (d *Downloader) DownloadAll() {
	videoExtractor := VideoExtractor{}

	for i := range d.EpList {
		episode := utils.EpisodeFormatter(i + 1)
		newDestPath := filepath.Join(d.DestPath, episode)

		if !CheckWebsite(d.EpList[i].URL) {
			message.ErrorMessage("Can't access episode url")
			return
		}

		videoExtractor.AnimeURL = d.EpList[i].URL
		videoUrl := videoExtractor.ExtractVideoURL()

		videoSrc := videoExtractor.ExtractActualVideoLabel(videoUrl)

		videoDownloader := VideoDownloader{DestPath: newDestPath, URL: videoSrc}
		videoDownloader.Download()
	}
}
