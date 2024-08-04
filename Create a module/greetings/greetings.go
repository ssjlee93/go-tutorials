package greetings

import (
	"errors"
	"fmt"
	"math/rand"
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
	message := fmt.Sprintf(randomFormat(), name)
	// nil = no error
	return message, nil
}

/*
@return one of the set msgs randomly
*/
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see yu, %v",
		"Hail, %v. Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
