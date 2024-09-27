package main

import (
	"fmt"
	"image"
)

func main() {
	// uses image package and its interfaces to create RGBA struct and Rectangle struct within RGBA
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// use image interface to call these methods
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
