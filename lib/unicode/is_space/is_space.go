package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune{' ', '\t', '\v', '\f', '\r', '🚀'}

	for _, v := range s {
		x := unicode.IsSpace(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
