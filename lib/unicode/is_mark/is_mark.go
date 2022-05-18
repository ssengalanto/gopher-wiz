package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "G̀ophèr Ẁiz"

	for _, v := range s {
		x := unicode.IsMark(v)
		fmt.Printf("%c: %t\n", v, x)
	}
}
