package main

import "fmt"

func hello(s string) string {
	return fmt.Sprintf("Hello %s!", s)
}

func main() {
	fmt.Println(hello("Gopher Wiz"))
}
