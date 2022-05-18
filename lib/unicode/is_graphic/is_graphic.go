package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 0xd (13) is unicode code point of enter
	x := unicode.IsGraphic(0xd)
	fmt.Println("x:", x)

	// 0x7e (126) is unicode point of '~'
	y := unicode.IsGraphic(0x7e)
	fmt.Println("y", y)
}
