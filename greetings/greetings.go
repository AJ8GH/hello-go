package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.sprintf("Hi, %v. Welcome!", name)
	return message
}
