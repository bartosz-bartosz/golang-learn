package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting
func Hello(name string) (string, error) {
	// If no name given - return error
	if name == "" {
		return "", errors.New("Error: no name provided")
	}
	// Return a greeting
	message := fmt.Sprintf(randomFormat(), name) // Valid call
	// message := fmt.Sprint(randomFormat()) // Test-failing call
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome.",
		"Great to see youu, %v!",
		"YOOO %v WASSUP???!",
	}

	// rand.Seed(time.Now().UnixNano())

	return formats[rand.Intn(len(formats))]
}
