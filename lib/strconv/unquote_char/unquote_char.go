package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x, mb, t, err := strconv.UnquoteChar(`\"Gopher's\"`, '"')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("x: %c, mb: %v, tail: %s", x, mb, t)
}
