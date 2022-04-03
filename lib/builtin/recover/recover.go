package main

import (
	"fmt"
)

func hello(name *string) {
	defer recoverFn()
	if name == nil {
		panic("runtime error: name cannot be nil")
	}
	fmt.Println(name)
}

func main() {
	hello(nil)
}

func recoverFn() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}
