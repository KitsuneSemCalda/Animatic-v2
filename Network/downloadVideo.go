package network

import (
	message "animatic-v2/Message"
	"fmt"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

type VideoDownloader struct {
	DestPath string
	URL      string
}

func (vd *VideoDownloader) Download() {
	client := grab.NewClient()

	outputPath := vd.DestPath + ".mp4"

	request, _ := grab.NewRequest(outputPath, vd.URL)
	resp := client.Do(request)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("%.2f%%\r",
				100*resp.Progress())
		case <-resp.Done:
			break loop
		}
	}

	if err := resp.Err(); err != nil {
		message.ErrorMessage(err.Error())
		return
	}

	message.SucessMessage(fmt.Sprintf("%s was downloaded to %s\n", vd.URL, vd.DestPath))
	return
}
