package main

import "fmt"

func main() {
	a := make([]int, 3)
	x := copy(a, []int{0, 1, 2, 3, 4, 5})
	fmt.Printf("%v elements copied. %v\n", x, a)

	b := make([]byte, 6)
	y := copy(b, "Gopher Wizard")
	fmt.Printf("%v elements copied. %v\n", y, string((b)))
}
