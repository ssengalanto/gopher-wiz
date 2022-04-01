package main

import "fmt"

func main() {
	a := make([]string, 7)
	fmt.Printf("%v\n", cap(a))

	b := make([]string, 7, 10)
	fmt.Printf("%v\n", cap(b))

	c := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", cap(c))

	d := make(chan string, 7)
	fmt.Printf("%v\n", cap(d))
}
