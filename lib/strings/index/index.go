package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	x := strings.Index(s, "Go")
	y := strings.Index(s, "Wiz")
	z := strings.Index(s, "go")

	fmt.Printf("%d\n%d\n%d", x, y, z)
}
