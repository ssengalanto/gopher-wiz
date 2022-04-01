package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4, 5}
	ints := append(a, b...)
	fmt.Println(ints)

	s := "Gopher"
	bytes := append([]byte(s), " Wizard"...)
	fmt.Println(string(bytes))
}
