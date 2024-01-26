package network

import (
	"net/http"
	"testing"

	"github.com/cavaliergopher/grab/v3"
)

type TDownloader interface {
	NewRequest(outputPath string, url string) error
	Do(request *grab.Request) *grab.Response
}

type Messenger interface {
	ErrorMessage(msg string)
	SuccessMessage(msg string)
}

type MockDownloader struct{}
type MockMessenger struct{}

func (d MockDownloader) NewRequest(outputPath string, url string) error {
	return nil
}

func (d MockDownloader) Do(request *grab.Request) *grab.Response {
	return &grab.Response{HTTPResponse: &http.Response{StatusCode: 200}}
}

func (m MockMessenger) ErrorMessage(msg string) {
}

func (m MockMessenger) SuccessMessage(msg string) {
}

func tdownloadVideo(d MockDownloader, m Messenger, destPath string, url string) {
}

func TestDownloadVideo(t *testing.T) {
	d := MockDownloader{}
	m := MockMessenger{}
	destPath := "/test/path"
	url := "http://testvideo.com"

	tdownloadVideo(d, m, destPath, url)
}
