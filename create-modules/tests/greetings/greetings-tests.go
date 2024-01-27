package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("Empty name")
	}
	// Create a message using a random format
	// Swap the following two lines to force the test fails
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

// randomFormat() returns one of a set of greeting messages
// The returned message is selected at random
// 📌 randomFormat() starts with a lowercase letter, making it accessible only in its own package (not exported)
func randomFormat() string {
	// A slice of message formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v. Well met!",
	}

	// Return a randomly selected message format by specifying a random index for slice of formats
	return formats[rand.Intn(len(formats))]
}
