package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(3)

	// Request a greeting
	message, err := greetings.Hello("Dawg")
	// If error was returned - print it and exit
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned print the returned message
	fmt.Println(message)

	names := []string{"Bob", "Alan", "Mary"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
