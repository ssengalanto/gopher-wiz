package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		a    int
		b    bool
		c, d string
	)

	n, err := fmt.Sscan("0 false Gopher Wiz", &a, &b, &c, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, "items scanned.", a, b, c, d)
}
