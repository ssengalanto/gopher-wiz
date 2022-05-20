package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'Α', 'Β', 'Γ', 'Δ', 'Ε', 'Ζ'}

	for _, v := range s {
		x := unicode.ToLower(v)
		fmt.Printf("%c: %c\n", v, x)
	}
}
