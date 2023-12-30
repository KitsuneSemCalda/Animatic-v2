package message

import "fmt"

// This function being to print a error message with style
func ErrorMessage(Message string) {
	fmt.Printf("\033[1;31m \u26A0 %s \033[m\n", Message)
	return
}
