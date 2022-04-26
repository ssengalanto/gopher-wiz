package main

import (
	"fmt"
	"strconv"
)

func main() {
	if x, err := strconv.Atoi("3"); err == nil {
		fmt.Printf("x: %d, type: %[1]T\n", x)
	}

	if y, err := strconv.Atoi("3a"); err != nil {
		fmt.Printf("y: %v, error: %v", y, err)
	}
}
