package more_types

import (
	"fmt"
	"strings"
)

// Structs, Struct Fields, Pointers to structs, Struct literals
type Vertex struct {
	X int
	Y int
}

// Struct literals
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	pv = &Vertex{1, 2} // has type *Vertex
)

// Range
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// Maps, Maps Literals
type VertexMaps struct {
	Lat, Long float64
}

var m map[string]VertexMaps

// Maps Literals
var mLiterals = map[string]VertexMaps{
	"Bell Labs": VertexMaps{
		40.68433, -74.39967,
	},
	"Google": VertexMaps{
		37.42202, -122.08408,
	},
}

func MoreTypes() {
	fmt.Println("More types : structs, slices, and maps")

	// Pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Structs
	fmt.Println(Vertex{1, 2})

	// Struct Fields
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Pointers to structs
	ps := &v
	ps.X = 1e9
	fmt.Println(v)

	// Struct literals
	fmt.Println(v1, pv, v2, v3)

	// Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices
	var sSlices []int = primes[1:4]
	fmt.Println(sSlices)

	// Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aa := names[0:2]
	b := names[1:3]
	fmt.Println(aa, b)

	b[0] = "XXX"
	fmt.Println(aa, b)
	fmt.Println(names)

	// Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sLiterals := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sLiterals)

	// Slice defaults
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// Slice length and capacity
	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 		Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// 		Extend its length.
	s = s[:4]
	printSlice(s)

	// 		Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Nil slices
	var sNil []int
	fmt.Println(sNil, len(sNil), cap(sNil))
	if sNil == nil {
		fmt.Println("nil!")
	}

	// Create a slice with make
	aMake := make([]int, 5)
	printSliceMake("a", aMake)

	bMake := make([]int, 0, 5)
	printSliceMake("b", bMake)

	c := bMake[:2]
	printSliceMake("c", c)

	d := c[2:5]
	printSliceMake("d", d)

	// Slices of slices
	// 		Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 		The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	var sAppend []int
	printSlice(sAppend)

	// 		append works on nil slices.
	sAppend = append(sAppend, 0)
	printSlice(sAppend)

	// 		The slice grows as needed.
	sAppend = append(sAppend, 1)
	printSlice(sAppend)

	// 		We can add more than one element at a time.
	sAppend = append(sAppend, 2, 3, 4)
	printSlice(sAppend)

	// Range
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Range continued
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// Maps
	m = make(map[string]VertexMaps)
	m["Bell Labs"] = VertexMaps{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

// Slice length and capacity
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Create a slice with make
func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
