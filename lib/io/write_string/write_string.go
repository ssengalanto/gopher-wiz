package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	b := new(bytes.Buffer)

	n, err := io.WriteString(b, "Gopher Wiz")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n%d bytes copied.\n", b, n)
}
