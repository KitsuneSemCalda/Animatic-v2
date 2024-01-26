package network

import (
	message "animatic-v2/Message"
	structure "animatic-v2/Structures"
	tui "animatic-v2/Tui"
	utils "animatic-v2/Utils"

	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AnimeSearcher struct {
	AnimeName string
}

func (as *AnimeSearcher) SearchAnimeInAnimeFire() (string, string, error) {
	animeName := utils.NameAnimeTreating(as.AnimeName)
	currentPageUrl := fmt.Sprintf("%s/pesquisar/%s", "https://www.animefire.net", animeName)

	if CheckWebsite(currentPageUrl) == false {
		message.ErrorMessage("The site is unavailable")
		os.Exit(1)
	}

	for {
		response, err := http.Get(currentPageUrl)

		if err != nil {
			message.ErrorMessage(err.Error())
			os.Exit(1)
		}

		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)

		if err != nil {
			message.ErrorMessage(err.Error())
			os.Exit(1)
		}

		animes := make([]structure.Anime, 0)

		doc.Find(".row.ml-1.mr-1 a").Each(func(i int, s *goquery.Selection) {
			anime := structure.Anime{
				Name: strings.TrimSpace(s.Text()),
				Url:  s.AttrOr("href", ""),
			}
			animeName = strings.TrimSpace(s.Text())

			animes = append(animes, anime)
		})

		// Sort the animes slice to return dub episodes first and leg episodes last
		sort.Slice(animes, func(i, j int) bool {
			return strings.Contains(animes[i].Name, "Dublado")
		})

		if len(animes) > 0 {
			index := tui.SelectAnimes(animes)
			selectedAnime := animes[index]

			return selectedAnime.Url, selectedAnime.Name, nil
		}

		nextPage, exists := doc.Find(".pagination .next a").Attr("href")

		if !exists || nextPage == "" {
			message.ErrorMessage("No anime found with the given name")
			os.Exit(1)
		}

		currentPageUrl = currentPageUrl + nextPage

		if CheckWebsite(currentPageUrl) == false {
			message.ErrorMessage("Can't search in next page")
			os.Exit(1)
		}
	}
}
