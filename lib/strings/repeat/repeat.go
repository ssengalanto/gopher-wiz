package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Go "

	x := strings.Repeat(s, 3)
	fmt.Println("x:", x)
}
