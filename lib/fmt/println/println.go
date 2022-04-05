package main

import (
	"fmt"
	"log"
)

func main() {
	s := "Gopher Wiz"

	n, err := fmt.Println("Hello", s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%d bytes written.", n)
}
