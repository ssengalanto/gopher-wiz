package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}

	fmt.Printf("One second is %d microseconds.\n", d.Microseconds())
}
