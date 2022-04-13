package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wizard"

	x := strings.ContainsRune(s, 'a')
	fmt.Println("x:", x)

	y := strings.ContainsRune(s, 97)
	fmt.Println("y:", y)

	z := strings.ContainsRune(s, '🚀')
	fmt.Println("z:", z)
}
