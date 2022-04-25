package main

import "fmt"

func fibonacci(done chan interface{}, c chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}

func main() {
	c := make(chan int)
	done := make(chan interface{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		done <- struct{}{}
	}()
	fibonacci(done, c)
}
