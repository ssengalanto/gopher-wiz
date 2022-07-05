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

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("(%6s) = %s\n", r, d.Round(r).String())
	}
}
