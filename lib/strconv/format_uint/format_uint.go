package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.FormatUint(3, 10)
	fmt.Printf("x: %s, type: %[1]T\n", x)

	y := strconv.FormatUint(20, 16)
	fmt.Printf("y: %s, type: %[1]T\n", y)
}
