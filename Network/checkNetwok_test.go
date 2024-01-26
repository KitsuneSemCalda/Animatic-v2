package network

import (
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestHasNetwork(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name     string
		mockResp httpmock.Responder
		want     bool
	}{
		{"Test 1: Successful request", httpmock.NewStringResponder(200, ""), true},
		{"Test 2: 404 error", httpmock.NewStringResponder(404, ""), false},
	}

	for i := 3; i <= 20; i++ {
		testName := "Test " + strconv.Itoa(i) + ": Random scenario"
		tests = append(tests, struct {
			name     string
			mockResp httpmock.Responder
			want     bool
		}{testName, httpmock.NewStringResponder(200, ""), true})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://www.google.com", tt.mockResp)

			if got := HasNetwork(); got != tt.want {
				t.Errorf("HasNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}
