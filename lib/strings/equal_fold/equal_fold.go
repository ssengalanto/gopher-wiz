package main

import (
	"fmt"
	"strings"
)

func main() {
	b := strings.EqualFold("Gopher Wiz", "GOPHER WIZ")
	fmt.Println(b)
}
