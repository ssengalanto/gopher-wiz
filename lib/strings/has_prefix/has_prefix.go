package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	x := strings.HasPrefix(s, "Go")
	y := strings.HasPrefix(s, "Gopher")
	z := strings.HasPrefix(s, "Wiz")

	fmt.Printf("%t\n%t\n%t", x, y, z)
}
