package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs is a method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// methods exist as part of type
	fmt.Println(v.Abs())
}
