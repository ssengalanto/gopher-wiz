package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.Quote("Gopher Wiz")
	fmt.Printf("x: %s, len: %d", x, len(x))
}
