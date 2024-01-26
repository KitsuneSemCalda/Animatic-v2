package utils

import (
	"testing"
)

func TestSplitAnimeName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "Attack on Titan  Season 1", "Attack on Titan"},
		{"Test 2", "Death Note  Season 1", "Death Note"},
		{"Test 3", "One Punch Man  Season 2", "One Punch Man"},
		{"Test 4", "My Hero Academia  Season 3", "My Hero Academia"},
		{"Test 5", "Naruto  Season 4", "Naruto"},
		{"Test 6", "Dragon Ball Z  Season 5", "Dragon Ball Z"},
		{"Test 7", "Fullmetal Alchemist  Season 1", "Fullmetal Alchemist"},
		{"Test 8", "Hunter x Hunter  Season 2", "Hunter x Hunter"},
		{"Test 9", "Fairy Tail  Season 3", "Fairy Tail"},
		{"Test 10", "One Piece  Season 4", "One Piece"},
		{"Test 11", "Tokyo Ghoul  Season 1", "Tokyo Ghoul"},
		{"Test 12", "Bleach  Season 2", "Bleach"},
		{"Test 13", "Cowboy Bebop  Season 1", "Cowboy Bebop"},
		{"Test 14", "Demon Slayer  Season 2", "Demon Slayer"},
		{"Test 15", "Sword Art Online  Season 3", "Sword Art Online"},
		{"Test 16", "Steins Gate  Season 1", "Steins Gate"},
		{"Test 17", "Code Geass  Season 2", "Code Geass"},
		{"Test 18", "Gintama  Season 3", "Gintama"},
		{"Test 19", "JoJo's Bizarre Adventure  Season 4", "JoJo's Bizarre Adventure"},
		{"Test 20", "Neon Genesis Evangelion  Season 1", "Neon Genesis Evangelion"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitAnimeName(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}
