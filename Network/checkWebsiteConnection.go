package network

import (
	"log"
	"net/http"
	"time"
)

const Timeout = 3 * time.Second

func CheckWebsite(url string) bool {
	client := http.Client{
		Timeout: Timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to fetch the website %s: %v", url, err)
		return false
	}
	defer resp.Body.Close()

	return true
}
