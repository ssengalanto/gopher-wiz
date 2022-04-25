package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Gopher Wiz")

	buf := make([]byte, 10)
	x, err := io.ReadAtLeast(r, buf, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%d bytes copied.\n", buf, x)

	sb := make([]byte, 4)
	y, err := io.ReadAtLeast(r, sb, 6)
	if err != nil {
		fmt.Printf("\nerror: %v\n%d bytes copied.\n", err, y)
	}

	lb := make([]byte, 32)
	z, err := io.ReadAtLeast(r, lb, 32)
	if err != nil {
		fmt.Printf("\nerror: %v\n%d bytes copied.\n", err, z)
	}
}
