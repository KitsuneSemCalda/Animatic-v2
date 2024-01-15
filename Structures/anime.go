package Structures

type Episode struct {
	Number string
	URL    string
}

type Anime struct {
	Name     string
	Url      string
	Episodes []Episode
}

func (a *Anime) AddEpisode(e Episode) {
	a.Episodes = append(a.Episodes, e)
}

func (a *Anime) TotalEpisodes() int {
	return len(a.Episodes)
}
