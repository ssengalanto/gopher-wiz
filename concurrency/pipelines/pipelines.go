package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for _, i := range integers {
				select {
				case <-done:
					return
				case c <- i:
				}
			}
		}()
		return c
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for i := range intStream {
				select {
				case <-done:
					return
				case c <- i * multiplier:
				}
			}
		}()
		return c
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for i := range intStream {
				select {
				case <-done:
					return
				case c <- i + additive:
				}
			}
		}()
		return c
	}

	done := make(chan interface{})
	defer close(done)

	stream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, stream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
}
