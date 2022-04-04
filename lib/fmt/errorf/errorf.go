package main

import "fmt"

func main() {
	a, b := "pineapple", "pizza"
	err := fmt.Errorf("%s should not be on %s", a, b)
	fmt.Println(err.Error())
}
