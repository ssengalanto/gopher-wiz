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

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%d bytes copied.", n)
}
