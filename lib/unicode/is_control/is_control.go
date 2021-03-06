package main

import (
	"fmt"
	"unicode"
)

/*
	key codes - https://bit.ly/3NilELT
*/

func main() {
	// 0xd (13) is unicode code point of enter
	x := unicode.IsControl(0xd)
	fmt.Println("x:", x)

	// 0x7e (126) is unicode point of '~'
	y := unicode.IsControl(0x7e)
	fmt.Println("y", y)
}
