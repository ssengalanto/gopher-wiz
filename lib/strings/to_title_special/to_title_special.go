package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	x := strings.ToTitleSpecial(unicode.TurkishCase, "ğöpher wiz")
	fmt.Println("x:", x)
}
