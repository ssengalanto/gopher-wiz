package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.Trim("-_Gopher Wiz_-", "-_")
	fmt.Println("x:", x)

	y := strings.Trim("-_Gopher-Wiz_-", "-_")
	fmt.Println("y:", y)
}
