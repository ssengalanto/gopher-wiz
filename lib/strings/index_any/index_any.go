package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	x := strings.IndexAny(s, "Go")
	y := strings.IndexAny(s, "go")
	z := strings.IndexAny(s, "JS")

	fmt.Printf("%d\n%d\n%d", x, y, z)
}
