package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.QuoteToASCII(`"Gopher	Wiz 🚀"`)
	fmt.Println("x:", x)
}
