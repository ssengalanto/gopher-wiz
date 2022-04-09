package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"

	x := strings.Contains(s, "Gopher")
	y := strings.Contains(s, "Golang")
	fmt.Printf("%t\n%t", x, y)
}
