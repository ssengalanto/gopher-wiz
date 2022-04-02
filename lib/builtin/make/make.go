package main

import "fmt"

func main() {
	a := make(map[string]bool)
	fmt.Printf("map: %v\n", a)

	b := make([]int, 7)
	fmt.Printf("slice: %v\n", b)

	c := make(chan bool)
	fmt.Printf("channel: %v\n", c)

	d := make(chan bool, 7)
	fmt.Printf("buffered channel: %v\n", d)
}
