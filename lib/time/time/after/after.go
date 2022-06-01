package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	d := time.Date(1992, 3, 8, 0, 0, 0, 0, time.UTC)

	x := t.After(d)
	fmt.Println("x:", x)

	y := t.After(t.Add(10 * time.Second))
	fmt.Println("y:", y)
}
