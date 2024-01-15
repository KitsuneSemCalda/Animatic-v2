package utils

import "strings"

func NameAnimeTreating(animeName string) string {
	loweredName := strings.ToLower(animeName)
	spacelessName := strings.ReplaceAll(loweredName, " ", "-")
	return spacelessName
}
