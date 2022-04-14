package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.Split("Gopher,Wiz", ",")
	fmt.Printf("x: %#v\n", x)

	y := strings.Split(" Gopher Wiz ", "")
	fmt.Printf("y: %#v", y)
}
