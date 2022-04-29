package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := make([]byte, 1)
	x := strconv.AppendUint(a, 3, 10)
	fmt.Printf("x: %s, type: %[1]T\n", x)

	b := make([]byte, 2)
	y := strconv.AppendUint(b, 20, 16)
	fmt.Printf("y: %s, type: %[1]T\n", y)
}
