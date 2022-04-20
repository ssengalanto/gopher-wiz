package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}

	x := strings.TrimLeftFunc("-0_Gopher Wiz_0-", f)
	fmt.Println("x:", x)
}
