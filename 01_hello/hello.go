// package get_started_with_go

// Declare a main package
// a package is a way to group functions, and it's made up of all the files in the same directory
package main

/*
import fml package
fmt = formatting text
one of standard library packages
*/
import "fmt"

import "rsc.io/quote"

// main function
// executes by default when you run the main package
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
