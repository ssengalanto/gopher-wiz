package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.FormatBool(true)
	fmt.Printf("x: %s, type: %[1]T\n", x)

	y := strconv.FormatBool(false)
	fmt.Printf("y: %s, type: %[1]T\n", y)
}
