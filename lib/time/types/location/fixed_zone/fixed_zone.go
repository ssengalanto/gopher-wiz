package main

import (
	"fmt"
	"time"
)

func main() {
	loc := time.FixedZone("UTC-8", -8*60*60)

	x := time.Date(1992, 3, 8, 0, 0, 0, 0, loc)
	fmt.Println("x:", x)
}
