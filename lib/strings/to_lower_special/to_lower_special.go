package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	x := strings.ToLowerSpecial(unicode.TurkishCase, "ĞÖPHER WİZ")
	fmt.Println("x:", x)
}
