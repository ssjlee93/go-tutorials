package main

import (
	"errors"
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 1 {
		return x, errors.New(ErrNegativeSqrt(x).Error())
	}
	z := 1.0
	for i := 0; i <= 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
