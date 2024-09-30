package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func Insert[T any](head *List[T], val T) *List[T] {
	newNode := &List[T]{val: val}
	newNode.next = head
	return newNode
}

func Print[T any](head *List[T]) {
	current := head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
	fmt.Println()
}

func main() {
	head := Insert[int](nil, 3)
	head = Insert[int](head, 2)
	head = Insert[int](head, 1)
	Print(head)

}
