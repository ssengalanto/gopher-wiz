package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"

	x := strings.HasPrefix(s, "Go")
	fmt.Println("x:", x)

	y := strings.HasPrefix(s, "Gopher")
	fmt.Println("y:", y)

	z := strings.HasPrefix(s, "Wiz")
	fmt.Println("z:", z)
}
