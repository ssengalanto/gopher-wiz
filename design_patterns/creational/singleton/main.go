package main

import (
	"fmt"
	"log"

	"github.com/ssengalanto/gopher-wiz/design_patterns/creational/singleton/singleton"
)

func main() {
	for i := 0; i < 20; i++ {
		go singleton.Instance()
	}

	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
}
