package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	b := new(bytes.Buffer)

	name := "Gopher Wiz"
	s := fmt.Sprintln("Hello", fmt.Sprint(name, "!"))

	n, err := io.WriteString(b, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes written. %q", n, b)
}
