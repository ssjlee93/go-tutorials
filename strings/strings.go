package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// Exercise 1
	//sample := []byte(sampler)

	// raw print
	fmt.Println("Println:")
	fmt.Println(sample)

	// byte-by-byte loop
	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	// hexadecimal format
	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	// quoted format
	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	// quoted and ASCII only
	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	// Exercise 2
	fmt.Println("Exercise 2")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Printf("\n")

	// UTF-8 and string literals
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
	fmt.Println()

	fmt.Printf("rune literal : ")
	runeLiteral := '⌘'
	fmt.Println(string(runeLiteral))
	fmt.Println()

	fmt.Println("Range Loop")
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	fmt.Println()

	// [Exercise: Put an invalid UTF-8 byte sequence into the string. (How?) What happens to the iterations of the loop?]
	const invalid = "\x01\xbd\xaa"
	for index, runeValue := range invalid {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	fmt.Println()

	fmt.Println("Libraries")
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
