package main

import (
	"golang.org/x/tour/pic"
)

// Pic creates an image with each value in the 2D slice as greyscale.
func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy) /* type declaration */
	for i := range ans {
		ans[i] = make([]uint8, dx) /* again the type? */
		for j := range ans[i] {
			ans[i][j] = uint8((i + j) / 2)
		}
	}
	return ans
}

// Show displaces the results of Pic.
// only works in online version for Tour of Go
func main() {
	pic.Show(Pic)
}
