package main

import (
	"fmt"
	"unicode"
)

func main() {
	x := unicode.IsLetter('a')
	fmt.Println("x:", x)

	y := unicode.IsLetter('~')
	fmt.Println("y", y)
}
