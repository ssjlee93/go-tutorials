package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFileWithoutDefer demonstrates a case without defer
func CopyFileWithoutDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	// if line below fails,
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	// files will never close on failure
	dst.Close()
	src.Close()
	return
}

// CopyFile opens two files and copies contents.
// demonstrates usage of `defer` statements
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// ensure files will always close
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// a demonstrates that `defer` statement's args are always evaluated when `defer` is evaluated
func a() {
	i := 0
	defer fmt.Println(i) // 0
	i++
	return
}

// b demonstrates that `defer` is LIFO
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
	// prints 3210
}

// c demonstrates that `defer` can chain with named return value
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func Panic() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	// if below `defer` function doesn't exist,
	// `f()` panics and `goruntime` crashes.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	// does not run since `panic` will stop executions
	// "g(4) behaves like a call to `panic`"
	// moves on to `defer` execution
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
