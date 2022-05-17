package main

import (
	"fmt"
	"unicode"
)

func main() {
	x := unicode.IsDigit('1')
	fmt.Println("x:", x)

	y := unicode.IsDigit('a')
	fmt.Println("y:", y)
}
