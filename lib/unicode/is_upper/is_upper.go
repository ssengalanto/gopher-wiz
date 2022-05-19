package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'Ω', 'ω', 'a', 'A', 'β'}

	for _, v := range s {
		x := unicode.IsUpper(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
