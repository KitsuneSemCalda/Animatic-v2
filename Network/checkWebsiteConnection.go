package network

import (
	"net/http"
	"time"
)

func CheckWebsite(url string) bool {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	_, err := client.Get(url)
	if err != nil {
		return false
	}

	return true
}
