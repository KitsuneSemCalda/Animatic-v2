package message

import "fmt"

// SucessMessage prints a success message with style.
// It takes a string as input and returns nothing.
func SucessMessage(message string) {
	if message == "" {
		return
	}
	const greenColor = "\033[1;32m"
	const resetColor = "\033[0m"
	const checkMark = "\u2705"

	fmt.Printf("%s%s %s %s\n", greenColor, checkMark, message, resetColor)
}
