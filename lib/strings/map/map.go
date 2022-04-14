package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	f := func(r rune) rune {
		if r == ' ' {
			return '_'
		}
		return r
	}

	x := strings.Map(f, s)
	fmt.Println("x:", x)
}
