package main

import (
	"fmt"
	"github.com/ssengalanto/gopher-wiz/design_patterns/creational/singleton/singleton"
)

func main() {

	x := singleton.New()
	fmt.Println("x:", x.Greet())

	y := singleton.New()
	fmt.Println("y:", y.Greet())

	fmt.Println("x and y are the same instance:", x == y)
}
