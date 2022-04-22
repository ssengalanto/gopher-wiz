package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Gopher Wiz")

	n, err := io.CopyN(os.Stdout, r, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d bytes copied.\n", n)
}
