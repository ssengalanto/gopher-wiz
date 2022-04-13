package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"

	x := strings.HasSuffix(s, "iz")
	fmt.Println("x:", x)

	y := strings.HasSuffix(s, "Wiz")
	fmt.Println("y:", y)

	z := strings.HasSuffix(s, "Go")
	fmt.Println("z:", z)
}
