package main

import (
	"fmt"
	"github.com/ssjlee93/go-tutorials/deferpanicandrecover"
	"github.com/ssjlee93/go-tutorials/slicesintro"
	"github.com/ssjlee93/go-tutorials/tour/basics"
	"github.com/ssjlee93/go-tutorials/tour/flowcontrol"
	"github.com/ssjlee93/go-tutorials/tour/moretypes"
)

func main() {
	fmt.Println("ssjlee93 Go Tutorials")

	fmt.Println("Tour of Go")
	fmt.Println("Basics")
	basics.Basics()
	fmt.Println()

	fmt.Println("Flowcontrol")
	flowcontrol.FlowControl()
	fmt.Println()

	fmt.Println("Moretypes")
	moretypes.MoreTypes()
	fmt.Println()

	fmt.Println("deferpanicandrecover")
	deferpanicandrecover.Panic()
	fmt.Println()

	fmt.Println("slicesintro")
	slicesintro.Slicesintro()
	fmt.Println()

}
