package main

import (
	"fmt"
	"time"
)

/*
	Done channel also known as a channel that signals termination
	to prevent goroutine leaks.
*/

func main() {
	fn := func(done chan interface{}) <-chan int {
		i := 0
		c := make(chan int)

		go func() {
			defer close(c)
			for {
				select {
				case <-done:
					return
				default:
					c <- i
				}
				i++
			}
		}()
		return c
	}

	done := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Nanosecond)
		close(done)
	}()

	stream := fn(done)
	for v := range stream {
		fmt.Printf("recieved: %d\n", v)
	}

	fmt.Println("Done receiving!")
}
