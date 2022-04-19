package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.TrimLeft("-_Gopher Wiz_-", "-_")
	fmt.Println("x:", x)

	y := strings.TrimLeft("-_Gopher-Wiz_-", "-_")
	fmt.Println("y:", y)
}
