package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'\t', '\n', '!', 'z', '🚀'}

	for _, v := range s {
		x := unicode.IsPrint(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
