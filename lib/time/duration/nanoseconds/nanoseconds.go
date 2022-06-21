package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1µs")
	if err != nil {
		panic(err)
	}

	fmt.Printf("One microsecond is %d nanoseconds.\n", d.Nanoseconds())
}
