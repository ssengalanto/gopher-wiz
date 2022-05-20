package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []int{unicode.UpperCase, unicode.LowerCase, unicode.TitleCase}

	for _, v := range s {
		x := unicode.To(v, 'ω')
		fmt.Printf("%d: %c\n", v, x)
	}
}
