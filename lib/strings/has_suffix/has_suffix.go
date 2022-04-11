package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	x := strings.HasSuffix(s, "iz")
	y := strings.HasSuffix(s, "Wiz")
	z := strings.HasSuffix(s, "Go")

	fmt.Printf("%t\n%t\n%t", x, y, z)
}
