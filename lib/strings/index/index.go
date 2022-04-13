package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"

	x := strings.Index(s, "Go")
	fmt.Println("x:", x)

	y := strings.Index(s, "Wiz")
	fmt.Println("y:", y)

	z := strings.Index(s, "go")
	fmt.Println("z:", z)
}
