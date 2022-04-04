package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	s := "Gopher Wiz"
	b := new(bytes.Buffer)

	n, err := fmt.Fprintln(b, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%d bytes written. %q", n, b)
}
