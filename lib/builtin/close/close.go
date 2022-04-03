package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	done := make(chan bool)

	for _, val := range s {
		c <- val
	}
	close(c)

	go func() {
		for i := 0; i < len(s)*2; i++ {
			val, ok := <-c
			if !ok {
				fmt.Println("received from closed channel", val, ok)
			} else {
				fmt.Println("received from channel", val, ok)
			}
		}
		done <- true
	}()

	<-done
}
