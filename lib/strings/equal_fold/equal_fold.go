package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.EqualFold("Gopher Wiz", "GOPHER WIZ")
	fmt.Println("x:", x)
}
