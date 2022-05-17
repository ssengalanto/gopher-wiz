package main

import (
	"fmt"
	"unicode"
)

func main() {
	rt := unicode.Greek

	x := unicode.In('α', rt)
	fmt.Println("x:", x)

	y := unicode.In('a', rt)
	fmt.Println("y:", y)
}
