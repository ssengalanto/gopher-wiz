package main

import "fmt"

func main() {
	a := 7 + 3i
	fmt.Printf("a: %v\n", a)

	b := complex(7, 3)
	fmt.Printf("b: %v\n", b)

	c := complex(-7, -3)
	fmt.Printf("c: %v\n", c)
}
