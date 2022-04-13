package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Gopher"
	b := "Wiz"

	// a < b
	x := strings.Compare(a, b)
	fmt.Println("x:", x)

	// a == b
	y := strings.Compare(a, a)
	fmt.Println("y:", y)

	// a > b
	z := strings.Compare(b, a)
	fmt.Println("z:", z)
}
