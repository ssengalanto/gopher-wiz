package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	if x, err := strconv.ParseFloat("3.1415926535", 64); err == nil {
		fmt.Printf("x: %d, type: %[1]T\n", x)
	}

	if y, err := strconv.ParseFloat(fmt.Sprintf("%e", math.MaxFloat64), 32); err != nil {
		fmt.Printf("y: %v, error: %v", y, err)
	}
}
