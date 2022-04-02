package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("value: %v, length: %d\n", m, len(m))

	delete(m, "b")
	fmt.Printf("value: %v, length: %d", m, len(m))
}
