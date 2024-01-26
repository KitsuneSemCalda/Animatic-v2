package network

import (
	"animatic-v2/Structures"
	"testing"
)

type TestDownloader interface {
	CheckWebsite(url string) bool
	ExtractVideoUrl(url string) string
	ExtractActualVideoLabel(url string) string
	DownloadVideo(destPath string, videoSrc string)
}

type TestMockDownloader struct{}

func (d MockDownloader) CheckWebsite(url string) bool {
	return true
}

func (d MockDownloader) ExtractVideoUrl(url string) string {
	return "http://testvideo.com"
}

func (d MockDownloader) ExtractActualVideoLabel(url string) string {
	return "Test Video"
}

func (d MockDownloader) DownloadVideo(destPath string, videoSrc string) {
}

func MockDownloadAll(d MockDownloader, destPath string, anime Structures.Anime, epList []Structures.Episode) {
}

func TestDownloadAll(t *testing.T) {
	d := MockDownloader{}
	destPath := "/test/path"
	anime := Structures.Anime{Name: "Test Anime"}
	epList := []Structures.Episode{{Number: "1", URL: "http://testepisode.com"}}

	MockDownloadAll(d, destPath, anime, epList)
}
