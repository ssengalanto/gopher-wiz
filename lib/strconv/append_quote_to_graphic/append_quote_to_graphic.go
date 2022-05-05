package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]byte, 12)
	b = strconv.AppendQuoteToGraphic(b, "\U0001f680")
	fmt.Printf("b: %s, type: %[1]T\n", b)
}
