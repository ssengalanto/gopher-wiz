package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	rx := strings.NewReader("Gopher")
	ry := strings.NewReader("Wiz")
	buf := make([]byte, 2)

	n, err := io.CopyBuffer(os.Stdout, rx, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d bytes copied.\n", n)

	m, err := io.CopyBuffer(os.Stdout, ry, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d bytes copied.", m)
}
