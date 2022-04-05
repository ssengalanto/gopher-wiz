package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var (
		a    int
		b    bool
		c, d string
	)

	r := strings.NewReader("7\n true gopher\n wiz excluded")

	n, err := fmt.Fscan(r, &a, &b, &c, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, "items scanned.", a, b, c, d)
}
