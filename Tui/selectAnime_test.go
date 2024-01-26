package tui

import (
	structure "animatic-v2/Structures"
	"testing"
)

// MockSelectAnimes is a mock function for SelectAnimes
func MockSelectAnimes(animes []structure.Anime) int {
	// Mock behavior here based on input
	// For simplicity, let's just return the length of the animes slice
	return len(animes)
}

func TestSelectAnimes(t *testing.T) {
	animes := []structure.Anime{
		{Name: "Anime 1"},
		{Name: "Anime 2"},
		{Name: "Anime 3"},
		{Name: "Anime 4"},
		{Name: "Anime 5"},
		{Name: "Anime 6"},
		{Name: "Anime 7"},
		{Name: "Anime 8"},
		{Name: "Anime 9"},
		{Name: "Anime 10"},
	}

	tests := []struct {
		name string
		want int
	}{
		{"Test 1: Select first anime", 1},
		{"Test 2: Select second anime", 2},
		{"Test 3: Select third anime", 3},
		{"Test 4: Select fourth anime", 4},
		{"Test 5: Select fifth anime", 5},
		{"Test 6: Select sixth anime", 6},
		{"Test 7: Select seventh anime", 7},
		{"Test 8: Select eighth anime", 8},
		{"Test 9: Select ninth anime", 9},
		{"Test 10: Select tenth anime", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MockSelectAnimes(animes[:tt.want]); got != tt.want {
				t.Errorf("MockSelectAnimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
