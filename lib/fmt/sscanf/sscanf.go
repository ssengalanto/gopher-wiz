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

	n, err := fmt.Sscanf("0 false Gopher Wiz", "%d %t %s %s", &a, &b, &c, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, "items scanned.", a, b, c, d)
}
