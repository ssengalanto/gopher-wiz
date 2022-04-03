package main

import (
	"fmt"
)

func main() {
	a := complex(7, 3)
	fmt.Println("The real part of", a, "is", real(a))

	b := complex(-7, -3)
	fmt.Println("The real part of", b, "is", real(b))
}
