package message

import "fmt"

// ErrorMessage prints an error message with style.
// It takes a string as input and returns nothing.
func ErrorMessage(message string) {
	if message == "" {
		return
	}

	const redColor = "\033[1;31m"
	const resetColor = "\033[0m"
	const warningSign = "\u26A0"

	fmt.Printf("%s %s %s %s\n", redColor, warningSign, message, resetColor)
}
