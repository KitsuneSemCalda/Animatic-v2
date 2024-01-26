package network

import (
	"net/http"
	"testing"

	"github.com/cavaliergopher/grab/v3"
)

type mtDownloader interface {
	NewRequest(outputPath string, url string) error
	Do(request *grab.Request) *grab.Response
}

type mtMessenger interface {
	ErrorMessage(msg string)
	SuccessMessage(msg string)
}

type mtMockDownloader struct{}
type mtMockMessenger struct{}

func (d MockDownloader) mtNewRequest(outputPath string, url string) error {
	return nil
}

func (d MockDownloader) mtDo(request *grab.Request) *grab.Response {
	return &grab.Response{HTTPResponse: &http.Response{StatusCode: 200}}
}

func (m MockMessenger) mtErrorMessage(msg string) {
}

func (m MockMessenger) mtSuccessMessage(msg string) {
}

func mtdownloadVideo(d Downloader, m Messenger, destPath string, url string) {
}

func MTestDownloadVideo(t *testing.T) {
	d := MockDownloader{}
	m := MockMessenger{}
	destPath := "/test/path"
	url := "http://testvideo.com"

	tdownloadVideo(d, m, destPath, url)
}
