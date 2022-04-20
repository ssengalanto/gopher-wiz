package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.ToValidUTF8("Gopher Wiz 🚀", "�")
	fmt.Println("x:", x)
}
