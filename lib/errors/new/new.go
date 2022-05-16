package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("gopher wiz has no mana")
	if err != nil {
		fmt.Println("error:", err)
	}
}
