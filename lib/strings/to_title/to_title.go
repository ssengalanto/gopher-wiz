package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.ToTitle("Gopher Wiz")
	fmt.Println("x:", x)

	y := strings.ToTitle("α β γ δ ε ζ")
	fmt.Println("y:", y)
}
