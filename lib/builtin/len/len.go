package main

import "fmt"

func main() {
	a := make([]string, 7)
	fmt.Printf("%v\n", len(a))

	b := "Gopher Wizard"
	fmt.Printf("%v\n", len(b))

	c := map[string]string{
		"key1": "value",
		"key2": "value",
	}
	fmt.Printf("%v\n", len(c))

	d := make(chan string)
	fmt.Printf("%v\n", len(d))
}
