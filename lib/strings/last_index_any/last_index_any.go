package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go - Gopher Wizard"

	x := strings.LastIndexAny(s, "Go")
	fmt.Println("x:", x)

	y := strings.LastIndexAny(s, "go")
	fmt.Println("y:", y)

	z := strings.LastIndexAny(s, "JS")
	fmt.Println("z:", z)
}
