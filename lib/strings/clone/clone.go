package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gopher Wiz"
	c := strings.Clone(s)

	fmt.Println(c)
}
