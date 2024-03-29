package main

import (
	"fmt"
	"log"
	"example.com/loops/greetings"
)

func main() {
	// Set the properties of the predefined logger, including:
	// - the log entry prefix and a flag to disable printing
	// - the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string {"Kyrex", "Will", "Joe"}

	// Request a greeting messages for the names
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of messages to the console
	fmt.Println(messages)
}
