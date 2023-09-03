package main

import (
	"fmt"
	"log"
	"io.github.aj8gh/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
