package main

/*
Go imports published Go modules
since our example is not published, we need to redirect the import to use local module
*/
import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Steve S. Lee")
	fmt.Println(message)
}
