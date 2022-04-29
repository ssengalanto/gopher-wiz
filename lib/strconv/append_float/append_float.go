package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := 3.1415926535
	s := []byte{'b', 'e', 'E', 'f', 'g', 'G', 'x', 'X'}

	for _, f := range s {
		b := make([]byte, 32)
		b = strconv.AppendFloat(b, v, f, -1, 64)
		fmt.Printf("value: %s, type: %[1]T\n", b)
	}
}
