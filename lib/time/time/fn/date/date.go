package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Date(1992, 3, 8, 0, 0, 0, 0, time.UTC)
	fmt.Println("x:", x)
}
