package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Gopher Wiz")

	sb := make([]byte, 6)
	n, err := io.ReadFull(r, sb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%d bytes copied.\n", sb, n)

	lb := make([]byte, 32)
	m, err := io.ReadFull(r, lb)
	if err != nil {
		fmt.Printf("\nerror: %v\n%d bytes copied.\n", err, m)
	}
}
