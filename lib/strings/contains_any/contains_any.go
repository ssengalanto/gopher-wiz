package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wizard"

	x := strings.ContainsAny(s, "Go")
	fmt.Println("x:", x)

	y := strings.ContainsAny(s, "Golang")
	fmt.Println("y:", y)

	z := strings.ContainsAny(s, "JS")
	fmt.Println("z:", z)
}
