package network

import (
	"fmt"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"gopkg.in/cheggaaa/pb.v1"
)

type VideoDownloader struct {
	DestPath string
	URL      string
}

func (vd *VideoDownloader) Download() {
	client := grab.NewClient()
	req, _ := grab.NewRequest(vd.DestPath+".mp4", vd.URL)
	req.HTTPRequest.Header.Set("Connection", "keep-alive")

	resp := client.Do(req)

	maxSizeInMB := int(resp.Size() / (1024 * 1024))
	minSizeInMB := 10

	if maxSizeInMB < minSizeInMB {
		maxSizeInMB = minSizeInMB
	}

	bar := pb.StartNew(maxSizeInMB)

	fmt.Printf("Episode URL: %s \n", vd.URL)
	for !resp.IsComplete() {
		completedInMB := int(resp.Progress() * float64(maxSizeInMB))
		bar.Set(completedInMB)
		time.Sleep(time.Millisecond * 500)
	}

	bar.Finish()
}
