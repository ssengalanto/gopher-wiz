package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'Ⅸ', '3', 'Ⅿ', 'M'}

	for _, v := range s {
		x := unicode.IsNumber(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
