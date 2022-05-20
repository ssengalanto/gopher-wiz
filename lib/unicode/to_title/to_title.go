package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'α', 'β', 'γ', 'δ', 'ε', 'ζ'}

	for _, v := range s {
		x := unicode.ToTitle(v)
		fmt.Printf("%c: %c\n", v, x)
	}
}
