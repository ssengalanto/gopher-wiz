package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wizard"

	x := strings.ContainsRune(s, 'a')
	y := strings.ContainsRune(s, 97)
	z := strings.ContainsRune(s, '🚀')
	fmt.Printf("%t\n%t\n%t", x, y, z)
}
