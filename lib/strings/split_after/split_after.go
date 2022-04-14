package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.SplitAfter("Gopher,Wiz,", ",")
	fmt.Printf("x: %#v\n", x)

	y := strings.SplitAfter(" Gopher Wiz ", "")
	fmt.Printf("y: %#v\n", y)

	z := strings.SplitAfter("", "")
	fmt.Printf("z: %#v", z)
}
