package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	x := strings.IndexRune(s, 'G')
	y := strings.IndexRune(s, 'o')
	z := strings.IndexRune(s, '🚀')

	fmt.Printf("%d\n%d\n%d", x, y, z)
}
