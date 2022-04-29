package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := make([]byte, 1)
	x := strconv.AppendInt(a, 3, 10)
	fmt.Printf("x: %s, type: %[1]T\n", x)

	b := make([]byte, 5)
	y := strconv.AppendInt(b, 20, 2)
	fmt.Printf("y: %s, type: %[1]T\n", y)
}
