package Structures

import (
	"encoding/json"
	"testing"
)

func TestVideoResponse(t *testing.T) {
	jsonStr := `{
		"data": [
			{
				"src": "http://testvideo.com/1",
				"label": "Test Video 1"
			},
			{
				"src": "http://testvideo.com/2",
				"label": "Test Video 2"
			}
		]
	}`

	var resp VideoResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 videos, got %d", len(resp.Data))
	}

	if resp.Data[0].Src != "http://testvideo.com/1" {
		t.Errorf("Expected first video src to be 'http://testvideo.com/1', got '%s'", resp.Data[0].Src)
	}

	if resp.Data[0].Label != "Test Video 1" {
		t.Errorf("Expected first video label to be 'Test Video 1', got '%s'", resp.Data[0].Label)
	}

	if resp.Data[1].Src != "http://testvideo.com/2" {
		t.Errorf("Expected second video src to be 'http://testvideo.com/2', got '%s'", resp.Data[1].Src)
	}

	if resp.Data[1].Label != "Test Video 2" {
		t.Errorf("Expected second video label to be 'Test Video 2', got '%s'", resp.Data[1].Label)
	}
}
