package network

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckWebsite(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	tests := []struct {
		name string
		url  string
		want bool
	}{
		{"Test 1: Valid URL", server.URL, true},
		{"Test 2: Invalid URL", "invalid_url", false},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckWebsite(tt.url); got != tt.want {
				t.Errorf("CheckWebsite() = %v, want %v", got, tt.want)
			}
		})
	}
}
