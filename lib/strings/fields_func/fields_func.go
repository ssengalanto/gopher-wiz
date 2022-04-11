package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "  Gopher🚀Wiz😱!😖"
	x := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("%#v\n", strings.FieldsFunc(s, x))
}
