package main

import (
	"fmt"
	"time"
)

func main() {
	s := "10h30m55s"
	x, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("x: %s, type: %[1]T", x)
}
