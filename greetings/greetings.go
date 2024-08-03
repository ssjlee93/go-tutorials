package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received,
	// Return a greeting that embeds the name in a message.
	/*
		Sprintf = string print format
		:= declares and initializes
		:= is the same as
			var message string
			message = fmt ...
	*/
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// nil = no error
	return message, nil
}
