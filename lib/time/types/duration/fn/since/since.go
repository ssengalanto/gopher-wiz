package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(1992, 3, 8, 0, 0, 0, 0, time.Local)
	x := time.Since(t)
	fmt.Printf("x: %.2f hrs", x.Hours())
}
