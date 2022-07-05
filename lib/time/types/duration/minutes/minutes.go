package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	h, m, s := t.Clock()
	duration := fmt.Sprintf("%dh%dm%ds", h, m, s)

	d, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}

	x := d.Minutes()
	fmt.Printf("x: %.2f mins", x)
}
