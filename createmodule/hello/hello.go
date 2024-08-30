package main

/*
Go imports published Go modules
since our example is not published, we need to redirect the import to use local module
*/
import (
	"bufio"
	"example.com/greetings"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Set properties of the predefined Logger
	// SetPrefix : the log entry prefix
	// SetFlags :flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	// added a functionality to read input from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	// request a greeting message
	// error case
	message, err := greetings.Hello(text)
	// If an error was returned,
	// log it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned,
	// print the message to the console
	fmt.Println(message)

	// ========== Hellos ==========

	// A slice of names.
	names := make([]string, 0)

	// logic to take input from console
	more := true

	for more {
		fmt.Println("Enter a name: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		names = append(names, text)

		fmt.Println("Do you have more? (y/n")
		ans, _ := reader.ReadString('\n')
		ans = strings.Replace(ans, "\n", "", -1)
		if ans == "n" {
			more = false
		} else if ans == "y" {
			continue
		} else {
			fmt.Println("Invalid input")
		}
	}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	for _, msg := range messages {
		fmt.Println(msg)
	}
}
