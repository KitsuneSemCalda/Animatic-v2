package network

import (
	"net/http"
	"time"
)

// HasNetwork checks if the system has network connectivity.
// It returns true if network is available, false otherwise.
func HasNetwork() bool {
	const url = "https://www.google.com"
	const timeout = 3 * time.Second

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()

	return true
}
