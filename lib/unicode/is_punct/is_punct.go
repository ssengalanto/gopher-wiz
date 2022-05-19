package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{'?', ';', '!', 'z', '🚀'}

	for _, v := range s {
		x := unicode.IsPunct(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
