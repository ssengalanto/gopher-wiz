package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]byte, 12)
	b = strconv.AppendQuoteToASCII(b, "Gopher Wiz")
	fmt.Printf("b: %s, type: %[1]T\n", b)
}
