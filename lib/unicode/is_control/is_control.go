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
	fmt.Println(unicode.IsControl(0xd))

	// 0x7e (126) is unicode point of '~'
	fmt.Println(unicode.IsControl(0x7e))
}
