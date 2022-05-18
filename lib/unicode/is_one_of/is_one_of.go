package main

import (
	"fmt"
	"unicode"
)

func main() {
	rt := []*unicode.RangeTable{unicode.Greek, unicode.Arabic}

	x := unicode.IsOneOf(rt, 'α')
	fmt.Println("x:", x)

	y := unicode.IsOneOf(rt, '؎')
	fmt.Println("y:", y)

	z := unicode.IsOneOf(rt, 'z')
	fmt.Println("z:", z)
}
