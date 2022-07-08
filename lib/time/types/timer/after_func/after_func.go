package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	timer := time.AfterFunc(time.Second*3, func() {
		fmt.Printf("AfterFunc: Elapsed time %.0fs\n", time.Since(start).Seconds())
	})
	defer timer.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	<-done

	fmt.Printf("Done: Elapsed time %.0fs\n", time.Since(start).Seconds())
}
