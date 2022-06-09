package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Now()
	fmt.Printf("x: %s\n isZero: %t\n", x, x.IsZero())

	var y time.Time
	fmt.Printf("\ny: %s\nisZero: %t", y, y.IsZero())
}
