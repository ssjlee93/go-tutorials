package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	/*
		Sprintf = string print format
		:= declares and initializes
		:= is the same as
			var message string
			message = fmt ...
	*/
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
