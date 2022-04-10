package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	log := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %t\n", s, sep, before, after, found)
	}

	log(s, "Go")
	log(s, "ph")
	log(s, "iz")
	log(s, "Golang")
}
