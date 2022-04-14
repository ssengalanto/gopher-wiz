package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.Replace("Go Go Go", "o", "opher", 2)
	fmt.Println("x:", x)

	y := strings.Replace("Javascript Javascript Javascript", "Javascript", "Go", -1)
	fmt.Println("y:", y)
}
