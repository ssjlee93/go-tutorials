// "Packages, variables, and functions"

// "Packages
package main

// "Imports"
// factored import statements.
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// "Variables"
var c, python, java bool

// "Variables with initializers"
var ii, jj int = 1, 2

// "Basic types"
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// "Constants"
const Pi = 3.14

// "Numeric constants"
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	// "Packages"
	fmt.Println("My favorite num is", rand.Intn(10))
	// "Imports"
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// "Exported names"
	fmt.Println(math.Pi)
	// "Functions" "Functions continued
	fmt.Println(add(3, 5))
	// "Multiple results"
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	// "Named return values"
	fmt.Println(split(17))
	// "Variables"
	var i int
	fmt.Println(i, c, python, java)
	// "Variables with initializers"
	var cc, pythonn, javaa = true, false, "no!"
	fmt.Println(ii, jj, cc, pythonn, javaa)
	// "Short variable declarations"
	var iii, jjj int = 1, 2
	k := 3
	ccc, pythonnn, javaaa := true, false, "no!"

	fmt.Println(iii, jjj, k, ccc, pythonnn, javaaa)
	// "Basic types"
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// "Zero values"
	var iiii int  // 0
	var f float64 // 0
	var bb bool   // false
	var s string  // ""
	fmt.Printf("%v %v %v %q\n", iiii, f, bb, s)
	// "Type conversions"
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ff)
	fmt.Println(x, y, z)
	// "Type inferences"
	v := 42 // change me!
	v1 := 42.5
	v2 := 42.5 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v is of type %T\n", v1)
	fmt.Printf("v is of type %T\n", v2)
	// "Constants"
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	// "Numeric constants"
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// "Exported names"
// exported names are capitalized
func Pizza() {
	fmt.Println(math.Pi)
}

// "Exported names"
// not exported
func pizza() {
	fmt.Println("pizza")
}

// "Functions" "Functions continued"
// multiple args
func add(x, y int) int {
	// (x int, y int) also works
	return x + y
}

// "Multiple results"
func swap(x, y string) (string, string) {
	return y, x
}

// "Named return values"
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

// "Numeric constants"
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
