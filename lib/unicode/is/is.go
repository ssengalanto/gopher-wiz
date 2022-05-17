package main

import (
	"fmt"
	"unicode"
)

func main() {
	rt := unicode.Greek

	x := unicode.Is(rt, 'α')
	fmt.Println("x:", x)

	y := unicode.Is(rt, 'a')
	fmt.Println("y:", y)
}
