package main

import "fmt"

/*
	Channel owner is responsible for the ff: (also known as lexical confinement)
	1. Instantiate the channel
	2. Perform writes, or pass ownership to another goroutine
	3. Close the channel
	4. Encapsulate the previous three things in this list and expose them
	   via a read-only channel.
*/

func main() {
	fn := func() <-chan int {
		n := 3
		c := make(chan int, n)
		go func() {
			defer close(c)
			for i := 1; i <= n; i++ {
				c <- i
			}
		}()

		return c
	}

	stream := fn()
	for v := range stream {
		fmt.Printf("recieved: %d\n", v)
	}

	fmt.Println("Done receiving!")
}
