package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x, err := strconv.QuotedPrefix("\"Gopher\" Wiz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("x:", x)

	if y, err := strconv.QuotedPrefix("Gopher Wiz"); err != nil {
		fmt.Printf("y: %q, err: %v", y, err)
	}
}
