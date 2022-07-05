package main

import (
	"fmt"
	"time"
)

func main() {
	// hours in a year
	h := 8760

	start := time.Date(1992, 3, 8, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	diff := now.Sub(start)
	fmt.Printf("diff:"+
		" %.2f years\n", diff.Hours()/float64(h))

}
