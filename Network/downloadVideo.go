package network

import (
	message "animatic-v2/Message"
	"fmt"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func downloadVideo(destPath string, url string) {
	client := grab.NewClient()

	outputPath := destPath + ".mp4"

	request, _ := grab.NewRequest(outputPath, url)
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

	message.SucessMessage(fmt.Sprintf("%s was downloaded to %s\n", url, destPath))
	return
}
