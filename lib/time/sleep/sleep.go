package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(3 * time.Second)
	end := time.Since(start) / time.Second

	fmt.Printf("slept for %d seconds", end)
}
