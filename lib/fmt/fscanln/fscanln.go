package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	var (
		a    int
		b    bool
		c, d string
	)

	s := "1 true gopher wiz \n 0 false awesome go"

	r := strings.NewReader(s)

	for {
		n, err := fmt.Fscanln(r, &a, &b, &c, &d)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n, "items scanned.", a, b, c, d)
	}
}
