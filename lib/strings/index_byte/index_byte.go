package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"

	x := strings.IndexByte(s, 'G')
	fmt.Println("x:", x)

	y := strings.IndexByte(s, 'o')
	fmt.Println("y:", y)

	z := strings.IndexByte(s, 'g')
	fmt.Println("z:", z)
}
