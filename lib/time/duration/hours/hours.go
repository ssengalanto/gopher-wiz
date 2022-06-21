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

	x := d.Hours()
	fmt.Printf("x: %.2f", x)
}
