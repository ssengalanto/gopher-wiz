package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'!', 'a', '3', 'Ⅸ', '☺', '🚀'}

	for _, v := range s {
		x := unicode.IsSymbol(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
