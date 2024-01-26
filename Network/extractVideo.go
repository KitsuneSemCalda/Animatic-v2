package network

import (
	message "animatic-v2/Message"
	"animatic-v2/Structures"
	"encoding/json"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type VideoExtractor struct {
	AnimeURL string
}

func (ve *VideoExtractor) ExtractVideoURL() string {

	if CheckWebsite(ve.AnimeURL) == false {
		message.ErrorMessage("Can't access video website")
		return ""
	}

	response, err := http.Get(ve.AnimeURL)

	if err != nil {
		message.ErrorMessage(err.Error())
		return ""
	}

	doc, _ := goquery.NewDocumentFromReader(response.Body)

	videoElements := doc.Find("video")

	if videoElements.Length() > 0 {
		oldDataVideo, _ := videoElements.Attr("data-video-src")
		return oldDataVideo
	} else {
		videoElements = doc.Find("div")
		if videoElements.Length() > 0 {
			oldDataVideo, _ := videoElements.Attr("data-video-src")
			return oldDataVideo
		}
	}

	message.ErrorMessage("Can't find video element in html")
	return ""
}

func (ve *VideoExtractor) ExtractActualVideoLabel(videoSrc string) string {
	var videoResponse Structures.VideoResponse
	response, err := http.Get(videoSrc)

	if err != nil {
		message.ErrorMessage(err.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		message.ErrorMessage("Status Code is different from StatusOK")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		message.ErrorMessage(err.Error())
		return ""
	}

	err = json.Unmarshal(body, &videoResponse)

	if err != nil {
		message.ErrorMessage(err.Error())
		return ""
	}

	if len(videoResponse.Data) == 0 {
		message.ErrorMessage("No Video Found")
		return ""
	}

	return videoResponse.Data[len(videoResponse.Data)-1].Src
}
