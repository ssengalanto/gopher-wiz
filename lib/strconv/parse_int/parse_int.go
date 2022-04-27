package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	if x, err := strconv.ParseInt("3", 10, 64); err == nil {
		fmt.Printf("x: %d, type: %[1]T\n", x)
	}

	if y, err := strconv.ParseInt(fmt.Sprintf("%d", math.MaxInt32+1), 10, 32); err != nil {
		fmt.Printf("y: %v, error: %v", y, err)
	}
}
