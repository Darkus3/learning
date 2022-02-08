package greetings

import (
  "fmt"
  "errors"
)

// Return a personalized greetings
func Hello(name string) (string, error) {
  if name == "" { return "", errors.New("Empty Name") }
  greet := fmt.Sprintf("Hi, %v. Welcome!", name)
  return greet, nil
}
