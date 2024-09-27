package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 1
	i := a[0]
	fmt.Println(i) // i == 1
	// a[2] == 0, the zero value of the int type
	fmt.Println(a[2] == 0)

}
