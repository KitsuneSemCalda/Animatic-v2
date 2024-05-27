package debug

import (
	"log"
	"strings"
)

func LogError(message string, err error) {
	message = strings.ToLower(message)
	log.Printf("Can't be execute %s because %v occured\n", message, err)
}

func LogInfo(message string) {
	message = strings.ToLower(message)
	log.Printf("Executing %s\n", message)
}
