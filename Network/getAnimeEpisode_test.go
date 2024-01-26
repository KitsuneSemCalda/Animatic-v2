package network

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type mttHTTPClient interface {
	Get(url string) (*http.Response, error)
}

type mttMessenger interface {
	ErrorMessage(msg string)
}

type mttMockHTTPClient struct{}
type mttMockMessenger struct{}

func (c mttMockHTTPClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("OK")),
	}, nil
}

func (m MockMessenger) mttErrorMessage(msg string) {
}

func mtextractVideoUrl(c mttMockHTTPClient, m mttMockMessenger, animeUrl string) string {
	return "video extracted"
}

func mtextractActualVideoLabel(c mttMockHTTPClient, m mttMockMessenger, videoSrc string) string {
	return "video label"
}

func TestExtractVideoUrl(t *testing.T) {
	c := mttMockHTTPClient{}
	m := mttMockMessenger{}
	animeUrl := "http://testanime.com"

	mtextractVideoUrl(c, m, animeUrl)
}

func TestExtractActualVideoLabel(t *testing.T) {
	c := mttMockHTTPClient{}
	m := mttMockMessenger{}
	videoSrc := "http://testvideo.com"

	mtextractActualVideoLabel(c, m, videoSrc)
}
