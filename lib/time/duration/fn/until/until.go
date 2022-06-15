package main

import (
	"fmt"
	"time"
)

func main() {
	// hours in a year
	h := 8760
	t := time.Date(2052, 3, 8, 0, 0, 0, 0, time.Local)

	x := time.Until(t)
	fmt.Printf("x: %.2f yrs", x.Hours()/float64(h))
}
