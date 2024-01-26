package network

import (
	"net/http"
	"testing"

	"github.com/cavaliergopher/grab/v3"
)

type mttDownloader interface {
	NewRequest(outputPath string, url string) error
	Do(request *grab.Request) *grab.Response
}

type mtttMessenger interface {
	ErrorMessage(msg string)
	SuccessMessage(msg string)
}

type mttMockDownloader struct{}
type mtttMockMessenger struct{}

func (d MockDownloader) mttNewRequest(outputPath string, url string) error {
	return nil
}

func (d MockDownloader) mttDo(request *grab.Request) *grab.Response {
	return &grab.Response{HTTPResponse: &http.Response{StatusCode: 200}}
}

func (m MockMessenger) mtttErrorMessage(msg string) {
}

func (m MockMessenger) mttSuccessMessage(msg string) {
}

func mttdownloadVideo(d MockDownloader, m Messenger, destPath string, url string) {
}

func mtTestDownloadVideo(t *testing.T) {
	d := MockDownloader{}
	m := MockMessenger{}
	destPath := "/test/path"
	url := "http://testvideo.com"

	mttdownloadVideo(d, m, destPath, url)
}
