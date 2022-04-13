package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "  Gopher🚀Wiz😱!😖"
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	x := strings.FieldsFunc(s, f)
	fmt.Printf("x: %#v\n", x)
}
