package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go - Gopher Wizard"

	x := strings.LastIndex(s, "Go")
	fmt.Println("x:", x)

	y := strings.LastIndex(s, "r")
	fmt.Println("y:", y)

	z := strings.LastIndex(s, "go")
	fmt.Println("z:", z)
}
