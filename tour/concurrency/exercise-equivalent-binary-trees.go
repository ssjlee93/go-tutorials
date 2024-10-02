package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Walking(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chLeft := make(chan int)
	chRight := make(chan int)

	go Walking(t1, chLeft)
	go Walking(t2, chRight)

	for {
		vLeft, okLeft := <-chLeft
		vRight, okRight := <-chRight

		if okLeft != okRight || vLeft != vRight {
			return false
		}

		if !okLeft {
			break
		}
	}
	return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Print("Yes!")
	} else {
		fmt.Print("No!")
	}
}
