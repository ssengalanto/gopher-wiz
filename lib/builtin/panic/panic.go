package main

import "fmt"

func hello(name *string) {
	if name == nil {
		panic("runtime error: name cannot be nil")
	}
	fmt.Println(name)
}

func main() {
	hello(nil)
}
