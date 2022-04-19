package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	f := func(r rune) bool {
		return !unicode.IsLetter(r)
	}
	x := strings.TrimFunc("-0_Gopher Wiz_0-", f)
	fmt.Println("x:", x)

	f = func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}
	y := strings.TrimFunc("-0_Gopher Wiz_0-", f)
	fmt.Println("y:", y)
}
