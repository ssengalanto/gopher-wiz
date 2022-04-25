package main

import "fmt"

/*
	Ad hoc confinement works but prefer using lexical confinement
*/

func main() {
	s := []int{1, 2, 3, 4}

	fn := func(c chan<- int) {
		defer close(c)

		for _, v := range s {
			c <- v
		}
	}

	c := make(chan int)
	go fn(c)

	for v := range c {
		fmt.Println(v)
	}
}
