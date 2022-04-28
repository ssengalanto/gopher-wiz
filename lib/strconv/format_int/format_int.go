package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.FormatInt(3, 10)
	fmt.Printf("x: %s, type: %[1]T\n", x)

	y := strconv.FormatInt(20, 2)
	fmt.Printf("y: %s, type: %[1]T\n", y)
}
