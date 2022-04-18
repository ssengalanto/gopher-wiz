package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	x := strings.ToUpperSpecial(unicode.TurkishCase, "ğöpher wiz")
	fmt.Println("x:", x)
}
