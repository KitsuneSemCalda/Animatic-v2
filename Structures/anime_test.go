package Structures

import (
	"testing"
)

func TestAnime(t *testing.T) {
	anime := Anime{Name: "Test Anime", Url: "http://testanime.com"}

	if anime.Name != "Test Anime" {
		t.Errorf("Expected anime name to be 'Test Anime', got '%s'", anime.Name)
	}

	if anime.Url != "http://testanime.com" {
		t.Errorf("Expected anime URL to be 'http://testanime.com', got '%s'", anime.Url)
	}

	if anime.TotalEpisodes() != 0 {
		t.Errorf("Expected total episodes to be 0, got '%d'", anime.TotalEpisodes())
	}

	anime.AddEpisode(Episode{Number: "1", URL: "http://testanime.com/1"})

	if anime.TotalEpisodes() != 1 {
		t.Errorf("Expected total episodes to be 1 after adding an episode, got '%d'", anime.TotalEpisodes())
	}

	if anime.Episodes[0].Number != "1" {
		t.Errorf("Expected first episode number to be '1', got '%s'", anime.Episodes[0].Number)
	}

	if anime.Episodes[0].URL != "http://testanime.com/1" {
		t.Errorf("Expected first episode URL to be 'http://testanime.com/1', got '%s'", anime.Episodes[0].URL)
	}
}
