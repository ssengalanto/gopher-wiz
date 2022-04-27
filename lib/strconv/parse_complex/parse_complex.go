package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x, err := strconv.ParseComplex("(7+3i)", 128)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("x: %v, type: %[1]T\n", x)

	if y, err := strconv.ParseComplex("-(7+3i)", 128); err != nil {
		fmt.Printf("y: %v, error: %v\n", y, err)
	}
}
