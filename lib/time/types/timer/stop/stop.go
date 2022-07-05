package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	<-done

	fmt.Printf("Elapsed time %.0fs", time.Since(start).Seconds())
}
