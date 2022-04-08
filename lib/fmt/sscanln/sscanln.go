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

	n, err := fmt.Sscanln("0 false Gopher \n Wiz", &a, &b, &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, "items scanned.", a, b, c, d)
}
