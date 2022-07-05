package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer := time.NewTimer(time.Second * 10)

	go func() {
		time.Sleep(time.Second * 2)
		timer.Reset(time.Second * 3)
	}()

	<-timer.C

	fmt.Printf("Elapsed time %.0fs", time.Since(start).Seconds())
}
