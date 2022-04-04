package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	s := "Gopher Wiz"
	b := new(bytes.Buffer)

	n, err := fmt.Fprintf(b, "Hello %s", s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%d bytes written. %s", n, b)
}
