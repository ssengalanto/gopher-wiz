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
	s := fmt.Sprintf("Hello %s!", name)

	n, err := io.WriteString(b, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, "bytes written.", b)
}
