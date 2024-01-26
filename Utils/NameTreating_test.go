package utils

import (
	"testing"
)

func TestNameAnimeTreating(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "Attack on Titan", "attack-on-titan"},
		{"Test 2", "Death Note", "death-note"},
		{"Test 3", "One Punch Man", "one-punch-man"},
		{"Test 4", "My Hero Academia", "my-hero-academia"},
		{"Test 5", "Naruto", "naruto"},
		{"Test 6", "Dragon Ball Z", "dragon-ball-z"},
		{"Test 7", "Fullmetal Alchemist", "fullmetal-alchemist"},
		{"Test 8", "Hunter x Hunter", "hunter-x-hunter"},
		{"Test 9", "Fairy Tail", "fairy-tail"},
		{"Test 10", "One Piece", "one-piece"},
		{"Test 11", "Tokyo Ghoul", "tokyo-ghoul"},
		{"Test 12", "Bleach", "bleach"},
		{"Test 13", "Cowboy Bebop", "cowboy-bebop"},
		{"Test 14", "Demon Slayer", "demon-slayer"},
		{"Test 15", "Sword Art Online", "sword-art-online"},
		{"Test 16", "Steins Gate", "steins-gate"},
		{"Test 17", "Code Geass", "code-geass"},
		{"Test 18", "Gintama", "gintama"},
		{"Test 19", "JoJo's Bizarre Adventure", "jojo's-bizarre-adventure"},
		{"Test 20", "Neon Genesis Evangelion", "neon-genesis-evangelion"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NameAnimeTreating(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}
