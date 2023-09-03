package main

import (
	"fmt"
	"log"
	"io.github.aj8gh/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	message, err := greetings.Hello("Gladys")
	handle(err)
	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	handle(err)
	fmt.Println(messages)
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
