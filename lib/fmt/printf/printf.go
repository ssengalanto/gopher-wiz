package main

import (
	"fmt"
	"log"
)

func main() {
	s := "Gopher Wiz"

	n, err := fmt.Printf("Hello %s", s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%d bytes written.", n)
}
