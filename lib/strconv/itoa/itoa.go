package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.Itoa(3)
	fmt.Printf("x: %q, type: %[1]T\n", x)
}
