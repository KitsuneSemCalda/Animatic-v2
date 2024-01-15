package utils

import "fmt"

func EpisodeFormatter(ep int) string {
	return fmt.Sprintf("S01E%02d", ep)
}
