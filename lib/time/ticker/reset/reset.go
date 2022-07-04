package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	// reset the ticker after 5 seconds and set interval to 2 seconds
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Reset(time.Second * 2)
	}()

	go func() {
		time.Sleep(15 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			h, m, s := t.Clock()
			fmt.Printf("Current time: %d:%d:%d\n", h, m, s)
		}
	}
}
