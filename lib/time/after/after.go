package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	select {
	// suppose to block forever
	case <-c:
		fmt.Println("blocking case")
	case <-time.After(3 * time.Second):
		fmt.Println("3 seconds passed. Time out!")
	}
}
