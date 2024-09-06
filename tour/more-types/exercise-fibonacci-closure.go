package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// initial values for fibonacci()
	first, second := 0, 1
	return func() int {
		// for each call of fibonacci, it returns result
		result := first
		// for each call, update newer values
		first, second = second, first+second
		return result
	}
}

func main() {
	// call fibonacci()
	f := fibonacci()
	for i := 0; i < 10; i++ {
		// for each iteration, call fibonacci()
		fmt.Println(f())
	}
}
