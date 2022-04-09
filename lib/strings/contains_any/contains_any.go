package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wizard"

	x := strings.ContainsAny(s, "Go")
	y := strings.ContainsAny(s, "Golang")
	z := strings.ContainsAny(s, "JS")
	fmt.Printf("%t\n%t\n%t", x, y, z)
}
