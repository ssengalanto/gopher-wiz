package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go - Gopher Wizard"

	x := strings.LastIndexByte(s, 'G')
	fmt.Println("x:", x)

	y := strings.LastIndexByte(s, 'o')
	fmt.Println("y:", y)

	z := strings.LastIndexByte(s, 'g')
	fmt.Println("z:", z)
}
