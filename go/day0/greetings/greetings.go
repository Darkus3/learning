package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Return a *single* personalized greetings
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	//greet := fmt.Sprintf(randomFormat(), name)
	greet := fmt.Sprint(randomFormat())
	return greet, nil
}

// Return *multiple* greetings message that associate to each name
func Hellos(names []string) (map[string]string, error) {
	// Create a map
	messages := make(map[string]string)
	// Iterate over names[]
	for _, name := range names {
		// get individual greets
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// Initialization the random seeds for number generator
// It is automatically executed by go at program startup
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Return one of pre-defined greets, randomly
func randomFormat() string {
	// Slice := "dynamic table"
	formats := []string{
		"Bonjour %v, bienvenue à vous!",
		"Hi %v, Welcome!",
		"Yo %v! Bien ou bien ?!",
	}

	// Return a random format
	return formats[rand.Intn(len(formats))]
}
