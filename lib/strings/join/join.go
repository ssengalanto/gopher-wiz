package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"Gopher", "Wiz"}

	x := strings.Join(s, " ")
	y := strings.Join(s, "-")
	z := strings.Join(s, "")

	fmt.Printf("%s\n%s\n%s", x, y, z)
}
