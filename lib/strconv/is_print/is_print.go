package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.IsPrint('g')
	fmt.Println("x:", x)

	y := strconv.IsPrint('🚀')
	fmt.Println("y:", y)

	z := strconv.IsPrint('\003')
	fmt.Println("z:", z)
}
