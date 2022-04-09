package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Gopher"
	b := "Wiz"

	fmt.Println("a < b:", strings.Compare(a, b))
	fmt.Println("a == a:", strings.Compare(a, a))
	fmt.Println("b > a:", strings.Compare(b, a))
}
