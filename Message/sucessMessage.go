package message

import "fmt"

func SucessMessage(Message string) {
	fmt.Printf("\033[1;32m\u2705 %s \033[m\n", Message)
	return
}
