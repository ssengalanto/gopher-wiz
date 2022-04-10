package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.Count("Gopher Wiz", "Go")
	y := strings.Count("Go Programming Language", "a")
	z := strings.Count("empty", "")

	fmt.Printf("%d\n%d\n%d", x, y, z)
}
