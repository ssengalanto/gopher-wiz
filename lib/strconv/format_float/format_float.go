package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := 3.1415926535
	s := []byte{'b', 'e', 'E', 'f', 'g', 'G', 'x', 'X'}

	for _, f := range s {
		x := strconv.FormatFloat(v, f, -1, 64)
		fmt.Printf("value: %v, type: %[1]T\n", x)
	}
}
