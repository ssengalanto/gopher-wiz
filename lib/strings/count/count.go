package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.Count("Gopher Wiz", "Go")
	fmt.Println("x:", x)

	y := strings.Count("Go Programming Language", "a")
	fmt.Println("y:", y)

	z := strings.Count("empty", "")
	fmt.Println("z:", z)
}
