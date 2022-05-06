package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.CanBackquote("Gopher Wiz")
	fmt.Println("x:", x)

	y := strconv.CanBackquote("`Gopher Wiz`")
	fmt.Println("y:", y)
}
