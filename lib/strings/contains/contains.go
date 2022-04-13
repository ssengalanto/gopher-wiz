package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"

	x := strings.Contains(s, "Gopher")
	fmt.Println("x:", x)

	y := strings.Contains(s, "Golang")
	fmt.Println("y:", y)
}
