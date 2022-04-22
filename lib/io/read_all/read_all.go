package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Gopher Wiz")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(b))

}
