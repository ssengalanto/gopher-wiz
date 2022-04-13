package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Greek, c)
	}

	x := strings.IndexFunc("Greek, α β γ", f)
	fmt.Println("x:", x)

	y := strings.IndexFunc("English, a b y", f)
	fmt.Println("y:", y)
}
