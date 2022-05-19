package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'Ω', 'ω', 'ǅ', 'A', 'β'}

	for _, v := range s {
		x := unicode.IsTitle(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
