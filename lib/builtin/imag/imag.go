package main

import (
	"fmt"
)

func main() {
	a := complex(7, 3)
	fmt.Println("The imaginary part of", a, "is", imag(a))

	b := complex(-7, -3)
	fmt.Println("The imaginary part of", b, "is", imag(b))
}
