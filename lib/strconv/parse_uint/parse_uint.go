package main

import (
	"fmt"
	"strconv"
)

func main() {
	if x, err := strconv.ParseUint("3", 10, 64); err == nil {
		fmt.Printf("x: %d, type: %[1]T\n", x)
	}

	if y, err := strconv.ParseUint("-1", 10, 32); err != nil {
		fmt.Printf("y: %v, error: %v", y, err)
	}
}
