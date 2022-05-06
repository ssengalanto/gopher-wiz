package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := strconv.IsGraphic('g')
	fmt.Println("x:", x)

	y := strconv.IsGraphic('🚀')
	fmt.Println("y:", y)

	z := strconv.IsGraphic('\003')
	fmt.Println("z:", z)
}
