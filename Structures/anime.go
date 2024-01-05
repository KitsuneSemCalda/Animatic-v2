package Structures

type Episode struct {
	Number int
	URL    string
}

type Anime struct {
	Name     string
	Url      string
	Episodes []Episode
}
