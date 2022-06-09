package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	x := t.Location()
	fmt.Println("x:", x)

	y := t.UTC().Location()
	fmt.Println("y:", y)
}
