package network

import (
	message "animatic-v2/Message"
	"animatic-v2/Structures"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Anime struct {
	URL string
}

func (a *Anime) GetEpisodes() []Structures.Episode {
	if CheckWebsite(a.URL) == false {
		message.ErrorMessage("Can't connect with website")
		return nil
	}

	resp, err := http.Get(a.URL)

	if err != nil {
		message.ErrorMessage("Failed to get anime details")
		os.Exit(1)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		message.ErrorMessage("Failed to parse anime details")
		os.Exit(1)
	}

	episodeContainer := doc.Find("a.lEp.epT.divNumEp.smallbox.px-2.mx-1.text-left.d-flex")

	episodes := make([]Structures.Episode, 0)

	episodeContainer.Each(func(i int, s *goquery.Selection) {
		episodeNum := s.Text()
		episodeURL, _ := s.Attr("href")

		episode := Structures.Episode{
			Number: episodeNum,
			URL:    episodeURL,
		}
		episodes = append(episodes, episode)
	})

	return episodes
}
