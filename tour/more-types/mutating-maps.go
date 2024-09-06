package main

import "fmt"

func main() {
	// create a map<string, int>
	m := make(map[string]int)

	// insert
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// update
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// delete
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// test that a key exists with 2-value assignment
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
