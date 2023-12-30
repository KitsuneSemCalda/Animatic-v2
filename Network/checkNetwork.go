package network

import (
	"net/http"
	"time"
)

func HasNetwork() bool {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	_, err := client.Get("https://www.google.com")
	if err != nil {
		return false
	}
	return true
}
