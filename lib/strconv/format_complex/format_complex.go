package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := complex(7, 3)
	s := []byte{'b', 'e', 'E', 'f', 'g', 'G', 'x', 'X'}

	for _, f := range s {
		x := strconv.FormatComplex(c, f, -1, 64)
		fmt.Printf("value: %v, type: %[1]T\n", x)
	}
}
