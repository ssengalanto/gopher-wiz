package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.QuoteToGraphic("Gopher Wiz \U0001f680")
	fmt.Println("x:", x)
}
