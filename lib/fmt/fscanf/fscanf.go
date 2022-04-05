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

	r := strings.NewReader("7 true gopher\nwiz excluded")

	n, err := fmt.Fscanf(r, "%d %t %s\n%s", &a, &b, &c, &d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, "items scanned.", a, b, c, d)
}
