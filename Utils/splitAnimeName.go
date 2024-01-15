package utils

import (
	"strings"
)

func SplitAnimeName(AnimeName string) string {
	parts := strings.Split(AnimeName, "  ")
	return parts[0]
}
