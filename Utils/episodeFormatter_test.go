package utils

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEpisodeFormatter(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tests := make([]struct {
		name     string
		input    int
		expected string
	}, 20)

	for i := range tests {
		tests[i].name = fmt.Sprintf("Test %d", i+1)
		tests[i].input = rand.Intn(41) - 20 // Generates a random number between -20 and 20
		tests[i].expected = fmt.Sprintf("S01E%02d", tests[i].input)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EpisodeFormatter(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}
