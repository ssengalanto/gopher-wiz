package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		defer w.Close()
		s := []string{"Gopher", " Wiz"}
		for _, v := range s {
			fmt.Fprint(w, v)
		}
	}()

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%d bytes copied.\n", n)
}
