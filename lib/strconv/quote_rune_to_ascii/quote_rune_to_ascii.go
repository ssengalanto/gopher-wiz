package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.QuoteRuneToASCII('🚀')
	fmt.Println("x:", x)
}
