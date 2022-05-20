package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'Ω', 'ω', 'a', 'A', 'β'}

	for _, v := range s {
		x := unicode.SimpleFold(v)
		fmt.Printf("%c: %#U\n", v, x)
	}
}
