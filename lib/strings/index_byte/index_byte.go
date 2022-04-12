package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	x := strings.IndexByte(s, 'G')
	y := strings.IndexByte(s, 'o')
	z := strings.IndexByte(s, 'g')

	fmt.Printf("%d\n%d\n%d", x, y, z)
}
