package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"

	x := strings.IndexAny(s, "Go")
	fmt.Println("x:", x)

	y := strings.IndexAny(s, "go")
	fmt.Println("y:", y)

	z := strings.IndexAny(s, "JS")
	fmt.Println("z:", z)
}
